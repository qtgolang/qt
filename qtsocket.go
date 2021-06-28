package qt

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
	"sync"
)
var mhandle int64
var applock sync.Mutex

const (
	收到数据 = 1
	客户进入 = 2
	客户离开 = 3
)

// Server .
//
//服务器端类型结构
type Server struct {
	// Connect .
	//
	//客户连接
	Connect func(server *Server, handle int64)
	// Exit .
	//
	//客户退出
	//tmpdata 为记录的临时数据
	Exit func(server *Server, handle int64, tmpdata *qtJson)
	// Arrival .
	//
	// 数据到达
	Arrival func(server *Server, handle int64, data []byte)
	// err .
	//
	//运行时的错误
	err string
	// network .
	//
	// 服务器
	network net.Listener
	// netconnectarr .
	//
	// 已连接的客户数组
	netconnectarr map[int64]interface{}

	//单次接收多少字节
	bit int
}

// SetBit .
//
//单次接收多少字节,默认4096
func (tcp *Server) SetBit(bit int) {
	//在qt.SocketServer 修改默认值
	tcp.bit = bit
}

// GetClientNum .
//
//获取已连接的客户端数量
func (tcp *Server) GetClientNum() int {
	return len(tcp.netconnectarr)
}

// GetClientHandle .
//
//获取全部客户端句柄
func (tcp *Server) GetClientHandle() []int64 {
	var arr []int64
	for k := range tcp.netconnectarr {
		arr = append(arr, k)
	}
	return arr
}

// SetTmpData .
//
// 设置临时记录数据 ，当连接被断开时，数据就被清空了
//
// 在连接被断开时，会最后一次收到临时数据的回调
func (tcp *Server) SetTmpData(handle int64, Name string, Data interface{}) bool {
	switch k := tcp.netconnectarr[handle].(type) {
	case map[string]interface{}:
		switch s := k["tmp"].(type) {
		case string:
			sjson := JSON()
			sjson.Untie(s)
			sjson.SetVal(Name, Data)
			k["tmp"] = sjson.ToString()
			tcp.netconnectarr[handle] = k
			return true
		}
	}
	return false
}

// GetTmpData .
//
// 获取临时记录数据 返回的是 *qtJson.Json
//
// 当连接被断开时，数据已经就被清空了
//
// 在连接被断开时，会最后一次收到临时数据的回调
func (tcp *Server) GetTmpData(handle int64) *qtJson {
	sjson := JSON()
	switch k := tcp.netconnectarr[handle].(type) {
	case map[string]interface{}:
		switch s := k["tmp"].(type) {
		case string:
			sjson.Untie(s)
		}
	}
	return sjson
}

// GetClientAddr .
//
//获取客户端地址
func (tcp *Server) GetClientAddr(handle int64) string {
	switch k := tcp.netconnectarr[handle].(type) {
	case map[string]interface{}:
		switch s := k["coon"].(type) {
		case net.Conn:
			return s.RemoteAddr().String()

		}
	default:
		fmt.Println(reflect.TypeOf(k), "GetClientAddr")
	}
	return ""
}

// Stat .
//
//服务器启动
func (tcp *Server) Stat(port interface{}) bool {
	Mport := ""
	switch s := port.(type) {
	case string:
		Mport = s
	case int:
		Mport = strconv.Itoa(s)
	case int8:
		Mport = strconv.Itoa(int(s))
	case int16:
		Mport = strconv.Itoa(int(s))
	case int32:
		Mport = strconv.Itoa(int(s))
	case int64:
		Mport = strconv.Itoa(int(s))
	case float32:
		Mport = strconv.Itoa(int(s))
	case float64:
		Mport = strconv.Itoa(int(s))
	}
	tcp.netconnectarr = (make(map[int64]interface{}))
	var err error
	tcp.network, err = net.Listen("tcp", ":"+Mport)
	if err != nil {
		tcp.err = err.Error()
		return false
	}
	tcp.run()
	return true
}

// Run .
//
// 开始监听 请在此之前调用 Stat
//
//本方法只是用于暂停程序，不让程序执行执行后面的代码
func (tcp *Server) run() {
	for {
		Conn, err := tcp.network.Accept()
		if err != nil {
			tcp.err = "accept error:" + err.Error()
			time.Sleep(time.Second * 10)
		} else {
			sd := handle(Conn.RemoteAddr().String())
			stmp := make(map[string]interface{})
			stmp["coon"] = Conn
			stmp["tmp"] = "{}"
			tcp.netconnectarr[sd] = stmp
			if tcp.Connect != nil {
				go tcp.advise(客户进入, sd, nil, 0)
			}
			go tcp.read(sd)
		}

	}
}
func (tcp *Server) read(sd int64) {
	switch k := tcp.netconnectarr[sd].(type) {
	case map[string]interface{}:
		switch v := k["coon"].(type) {
		case net.Conn:
			for {
				t1 := time.Now()
				b := make([]byte, tcp.bit)
				n, err := v.Read(b)
				if err == nil {
					t2 := time.Now()
					t3 := t2.Sub(t1).Nanoseconds()
					if t3 < 100000 && n == 0 {
						sl := tcp.GetTmpData(sd).ToString()
						delete(tcp.netconnectarr, sd) //从记录的Map中删除
						if tcp.Exit != nil {
							go tcp.advise(客户离开, sd, []byte(sl), 0) //发送给用户，客户断开的通知
						}
						return //不在继续监听获取数据
					}
					if tcp.Arrival != nil {
						go tcp.advise(收到数据, sd, b, n) //发送给用户，客户断开的通知
					}
				} else if err.Error() == "EOF" || strings.Contains(err.Error(), "closed") {
					sl := tcp.GetTmpData(sd).ToString()
					delete(tcp.netconnectarr, sd) //从记录的Map中删除
					if tcp.Exit != nil {
						go tcp.advise(客户离开, sd, []byte(sl), 0) //发送给用户，客户断开的通知
					}
					return //不在继续监听获取数据
				} else {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
func (tcp *Server) advise(Type int, handles int64, data []byte, Len int) {
	switch Type {
	case 收到数据:
		if len(data) <= Len {
			tcp.Arrival(tcp, handles, data) //发送给用户，客户断开的通知
		} else {
			tcp.Arrival(tcp, handles, data[:Len]) //发送给用户，客户断开的通知
		}
	case 客户进入:
		tcp.Connect(tcp, handles) //发送给用户，客户进入的通知
	case 客户离开:
		stmp := JSON()
		stmp.Untie(string(data))
		tcp.Exit(tcp, handles, stmp) //发送给用户，客户断开的通知
	}
}
func handle(a string) int64 {
	tmp := Md5(a, nil)
	if len(tmp) != 16 {
		return 0
	}
	return abs(int64(binary.BigEndian.Uint64(tmp)))
}
func abs(a int64) (ret int64) {
	ret = (a ^ a>>31) - a>>31
	return
}

// CloseClient .
//
//关闭客户端句柄
func (tcp *Server) CloseClient(handle int64) {
	switch k := tcp.netconnectarr[handle].(type) {
	case map[string]interface{}:
		switch s := k["coon"].(type) {
		case net.Conn:
			s.Close()
		}
	}
}

// Send .
//
// @handle 客户句柄
//
// 发送数据给客户端
func (tcp *Server) Send(handle int64, data interface{}) bool {
	bytes := []byte{}
	switch da := data.(type) {
	case string:
		bytes = []byte(da)
	case []byte:
		bytes = da
	}
	if len(bytes) < 1 {
		return false
	}
	switch k := tcp.netconnectarr[handle].(type) {
	case map[string]interface{}:
		switch s := k["coon"].(type) {
		case net.Conn:
			n, e := s.Write(bytes)
			if n != len(bytes) {
				if e != nil {
					fmt.Print(e.Error(), "err")
				}
				fmt.Print("N", n)
			}
			return n == len(bytes)
		}
	}
	return false
}

// Close .
//
//服务器关闭
func (tcp *Server) Close(port string) {
	if tcp.network != nil {
		tcp.network.Close()
	}
}

// Client .
//
// 客户端类型结构
type Client struct {
	// synchronization .
	//
	// 是否为异步请求
	asynchronous    func(handle int64, data []byte)
	conn            net.Conn
	err             string
	sendTimesOut    time.Duration
	receiveTimesOut time.Duration
	connectTimesOut time.Duration
	handleID        int64
}

// SetAsynchronous .
//
// 是否为异步模式 ,如果不为 nil 则异步 默认同步模式
//
//在 Connect 之前调用
//
//返回一个 句柄 多线程可用来区分
func (tcp *Client) SetAsynchronous(SetAsynchronous func(handle int64, data []byte)) int64 {
	tcp.asynchronous = SetAsynchronous
	applock.Lock()
	mhandle++
	tcp.handleID = mhandle
	applock.Unlock() 
	return tcp.handleID
}

// Connect .
//
//连接到服务器
//
//默认同步,如需异步 请在 Connect 之前 调用 SetAsynchronous
func (tcp *Client) Connect(ServerIP, port string) bool {
	var err error
	if tcp.connectTimesOut == 0 {
		tcp.connectTimesOut = 15 * time.Second
	}
	tcp.conn, err = net.DialTimeout("tcp4", ServerIP+":"+port, tcp.connectTimesOut)
	if err != nil {
		tcp.err = err.Error()
		return false
	}
	go tcp.receive()
	return true
}
func (tcp *Client) receive() {
	if tcp.asynchronous != nil {
		for {
			tcp.asynchronous(tcp.handleID, tcp.Receive())
		}
	}
}

// Receive .
//
// 接收数据
func (tcp *Client) Receive() []byte {
	buffer := make([]byte, 4096)
	if tcp.asynchronous != nil {
		tcp.receiveTimesOut = 1 * time.Second
	} else if tcp.receiveTimesOut == 0 {
		tcp.receiveTimesOut = 15 * time.Second
	}
	tcp.conn.SetReadDeadline(time.Now().Add(tcp.receiveTimesOut))
	n, err := tcp.conn.Read(buffer)
	if err != nil {
		tcp.err = err.Error()
		return nil
	}
	return buffer[:n]
}

// GetErr .
//
// 获取错误信息
func (tcp *Client) GetErr() string {
	return tcp.err
}

// ClientClose .
//
// 断开连接
func (tcp *Client) ClientClose() {
	err := tcp.conn.Close()
	if err != nil {
		tcp.err = err.Error()
	}
}

// SendClient .
//
// 发送数据
//
//每次发送数据，等5000微秒，防止毡包
func (tcp *Client) ClientSend(Data interface{}) {
	//time.Sleep(time.Microsecond * 5000)
	if tcp.sendTimesOut == 0 {
		tcp.sendTimesOut = 15 * time.Second
	}
	err := tcp.conn.SetWriteDeadline(time.Now().Add(tcp.sendTimesOut))
	if err != nil {
		tcp.err = err.Error()
		return
	}
	switch v := Data.(type) {
	case string:
		if n, err := tcp.conn.Write([]byte(v)); err != nil {
			n = n + 1
			tcp.err = err.Error()
		}
	case []byte:
		if n, err := tcp.conn.Write(v); err != nil {
			n = n + 1
			tcp.err = err.Error()
		}

	}

}

// SetTimesOut .
//
// 设置超时,单位毫秒
//
// Send  发送数据超时
//
// Receive 接收数据超时
//
// connect 连接超时
func (tcp *Client) SetTimesOut(Send, Receive, connect time.Duration) {
	tcp.sendTimesOut = Send * time.Millisecond
	tcp.receiveTimesOut = Receive * time.Millisecond
	tcp.connectTimesOut = connect * time.Millisecond
}
