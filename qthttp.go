package qt

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"
)

// Winhttp .
//
//Winhttp结构体
type winhttp struct {
	setRedirect  bool //是否允许自动重定向
	timesConnect time.Duration
	timesSend    time.Duration
	timesReceive time.Duration
	proxyip      string //代理IP
	method       string //提交方法
	url          string
	gzip         bool //是否自动Gzip压缩解压
	head         map[string]string
	head2        map[string]string
	cookies      map[string]string //内部自动合并更新
	ret          preturnData
}

type preturnData struct {
	ret  []byte              //返回的数据
	err  string              //过程中遇到的错误
	code int                 //返回的状态码
	Head map[string][]string //返回的Head
}

// SetLocation 是否允许自动重定向，默认允许自动重定向
//
//请在Send之前调用
func (qt *winhttp) SetLocation(Redirect bool) {
	qt.setRedirect = Redirect
}

// GetErr 获取运行过程中捕获到的错误
func (qt *winhttp) GetErr() string {
	return qt.ret.err
}

// SetGzip 是否允许自动Gzip压缩解压
//
//如果开启 发送数据时 不会 自动添加 Accept-Encoding: qtgzip 协议头
//
//请在Send之前调用
func (qt *winhttp) SetGzip(compress bool) {
	qt.gzip = compress
}

// SetProxyIP 设置或取消代理IP,IP为空则取消代理，user，pass 如果没有则填空文本
//
//请在Send之前调用
func (qt *winhttp) SetProxyIP(ip string) {
	qt.proxyip = ip
}

// SetProxyBasicAuth 设置代理IP 账号密码 账号密码为空则取消
//
//请在Send之前调用
func (qt *winhttp) SetProxyBasicAuth(User, Pass string) {
	if User == "" && Pass == "" {
		delete(qt.head, "Authorization")
		delete(qt.head, "Proxy-Authorization")
		return
	}
	qt.head2["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(User+":"+Pass))
	qt.head2["Proxy-Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(User+":"+Pass))
}

// SetUserAgent 设置User-Agent
//
//请在Send之前调用
func (qt *winhttp) SetUserAgent(val string) {
	qt.head2["User-Agent"] = val
}

// GetText 获取返回数据的字符串 如果是gbk，则是GBK 如果是UTF8,则是UTF8
//
//请在Send之后调用
func (qt *winhttp) GetText() string {
	return string(qt.ret.ret)
}

// GetBody 返回原始二进制数据
//
//请在Send之后调用
func (qt *winhttp) GetBody() []byte {
	return qt.ret.ret
}
func (qt *winhttp) InitHead() {
	qt.head = make(map[string]string)
}

//GetBodyAuto 自动判定是否UTF-8 如果不是UTF8编码，自动转为UTF8编码
//
//请在Send之后调用
func (qt *winhttp) GetBodyAuto() string {
	str := strings.ToUpper(qt.GetResponseHeaderALL())
	Mstr := strings.ToUpper(qt.GetText())
	M := ""
	if strings.Contains(str+Mstr, "=GB23") {
		M = "GBK"
	} else if strings.Contains(str+Mstr, "=GBK") {
		M = "GBK"
	} else if strings.Contains(str+Mstr, "UTF-8") {
		M = "UTF-8"
	} else {
		if len(qt.ret.ret) > 3 {
			if qt.ret.ret[0] == 239 && qt.ret.ret[1] == 187 && qt.ret.ret[2] == 191 {
				M = "UTF-8"
			} else {
				M = "GBK"
			}
		} else {
			M = "GBK"
		}
	}
	if M == "GBK" {
		return useNewEncoder(string(qt.ret.ret), "gbk", "UTF-8")
	}

	enc := mahonia.NewEncoder(M)
	return enc.ConvertString(string(qt.ret.ret))
}
func useNewEncoder(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// SetHeader 设置要提交的协议头 HeaderValue 若为空 则 表示删除 HeaderValue 协议头。
//
//请在Send之前调用
//
//Send之后本次设置的Head将会被清空
func (qt *winhttp) SetHeader(Name, Val string) {
	if Name == "" {
		return
	} else if Val == "" {
		delete(qt.head, Name)
		return
	}
	qt.head[Name] = Val
}

// AddCookies 设置单条Cookies  name 若为空 则 表示删除 name Cookie。
//
//请在Send之前调用
func (qt *winhttp) AddCookies(name, Val string) {
	if name == "" {
		return
	} else if Val == "" {
		delete(qt.cookies, name)
		return
	}
	qt.cookies[name] = Val
}

// Cookies 如果参数不是空文本 则 Cookies合并更新，如果是空文本，则清空全部Cookies
//
//传入参数例如：PSTM=1555142153; delPer=0; BD_HOME=0; BD_UPN=1126314751
func (qt *winhttp) Cookies(cookies string) {
	if cookies == "" {
		qt.cookies = make(map[string]string)
	} else {
		arr1 := strings.Split(cookies, ";")
		for i := 0; i < len(arr1); i++ {
			arr2 := strings.Split(arr1[i], "=")
			if len(arr2) == 2 {
				qt.cookies[strings.Trim(arr2[0], " ")] = strings.Trim(arr2[1], " ")
			}
		}
	}

}

// GetCookie 如果参数不是空文本 取单条Cookie，不附带Cookies昵称
//
//如果参数等于空，则取全部Cookies
func (qt *winhttp) GetCookie(name string) string {
	if name != "" {
		return qt.cookies[name]
	}
	Str := ""
	for k, v := range qt.cookies {
		Str = Str + k + "=" + v + "; "
	}
	return Str
}

// GetStatus 获取返回的状态码
//
//请在Send之后调用
func (qt *winhttp) GetStatus() int {
	return qt.ret.code
}

// SetTimeouts 设置超时设置单位秒,不能小于1秒
//
// Connect 连接超时
//
// Send    发送超时
//
// Receive 相应超时
//
//请在Send之前调用
func (qt *winhttp) SetTimeouts(Connect, Send, Receive time.Duration) {
	if Connect < 1*time.Second {
		qt.timesConnect = 1 * time.Second
	} else {
		qt.timesConnect = Connect
	}
	//=====================================
	if Send < 1*time.Second {
		qt.timesSend = 1 * time.Second
	} else {
		qt.timesSend = Send
	}
	//=====================================
	if Receive < 1*time.Second {
		qt.timesReceive = 1 * time.Second
	} else {
		qt.timesReceive = Receive
	}
}

// Open 设置要提交的类型（POST or GET ...） 以及URL
//
//请在Send之前调用
func (qt *winhttp) Open(method, url string) {
	if method != "OPTIONS" && method != "GET" && method != "HEAD" && method != "PUT" && method != "DELETE" && method != "TRACE" && method != "CONNECT" && method != "POST" {
		method = "GET"
	}
	qt.method = strings.ToUpper(method)
	qt.url = url
}

// GetResponseHeaderALL 获取返回的全部协议头
//
//请在Send之后调用
func (qt *winhttp) GetResponseHeaderALL() string {
	str := ""
	for k, v := range qt.ret.Head {
		for l := 0; l < len(v); l++ {
			str = str + k + ": " + v[l] + "\r\n"
		}
	}
	return strings.Trim(str, "\r\n")
}

// GetResponseHeader 获取返回的单个协议头 ,将不返回协议头名称
//
//如果 同一个name 有多条，将返回多行
//
//例如：Set-Cookie
//
//Set-Cookie: MUID=33A958A57F7A605B35A955987B7A61B9; domain=.bing.com; expires=Thu, 07-May-2020 08:59:15 GMT; path=/;
//
//Set-Cookie: MR=0; domain=c.bing.com; expires=Thu, 10-Oct-2019 08:59:15 GMT; path=/;
//
//Set-Cookie: SRM_B=33A958A57F7A605B35A955987B7A61B9; domain=c.bing.com; expires=Thu, 07-May-2020 08:59:15 GMT; path=/;
//
//请在Send之后调用
func (qt *winhttp) GetResponseHeader(name string) string {
	str := ""
	for k, v := range qt.ret.Head {
		if k == name {
			for l := 0; l < len(v); l++ {
				str = str + v[l] + "\r\n"
			}
		}
	}
	return strings.Trim(str, "\r\n")
}

// Send .
// 开始访问,并且发送 字符串数据或二进制数据，如果没有要提交的数据请填空文本 或 空字节
func (qt *winhttp) Send(Data interface{}) {
	switch val := Data.(type) {
	case string:
		qt.sendBin([]byte(val))
	case []byte:
		qt.sendBin((val))
	}
}

// sendBin .
// 开始访问
func (qt *winhttp) sendBin(bin []byte) {
	qt.ret = *new(preturnData) //重新初始化 返回值需要的类型
	if qt.url == "" {
		qt.ret.err = "[winhttp][err] url nil"
		return
	}
	ContentType := false
	compress := false
	databin := bin
	for k, v := range qt.head {
		if strings.Contains(strings.ToLower(k), "content-type") {
			ContentType = true //在设置的所有协议头中寻找，Content-Type 这个协议头 ，因为如果是Post 请求，必须带上这个协议头
		}
		if strings.Contains(strings.ToLower(k), "accept-encoding") {
			if strings.Contains(strings.ToLower(v), "qtgzip") {
				compress = true //在设置的所有协议头中寻找，Accept-Encoding 这个协议头 如果有，Gzip压缩后再提交
			}
		}
	}
	if !ContentType && (qt.method == "PUT" || qt.method == "POST") {
		//如果协议头中没有这个，且 类型为 PUT 或者 POST 则自动给加上
		qt.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	if compress && qt.gzip && len(bin) > 1 {
		//如果协议头有 Accept-Encoding 这个协议头 Gzip压缩后再提交
		databin = NewGzip().Compress(bin)
	}

	req, err := http.NewRequest(qt.method, qt.url, bytes.NewBuffer(databin))
	if err != nil {
		qt.ret.err = "[winhttp][ NewRequest err]" + err.Error()
		return
	}
	for k, v := range qt.head2 {
		textproto.MIMEHeader(req.Header)[k] = []string{v}
		//req.Header.Set(k, v)
	}
	for k, v := range qt.head {
		textproto.MIMEHeader(req.Header)[k] = []string{v}
		//h[CanonicalMIMEHeaderKey(key)] = []string{value}
		//req.Header.Set(k, v)
	}
	qt.head = make(map[string]string) //清空用户设置的协议头
	cookie := qt.GetCookie("")
	if cookie != "" {
		//将Cookie添加置协议头中
		req.Header.Set("Cookie", cookie)
	}
	transport := &http.Transport{}

	Dial := func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, qt.timesConnect) //设置建立连接超时
		if err != nil {
			qt.ret.err = "[winhttp][ DialTimeout err]" + err.Error()
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(qt.timesSend)) //设置发送接受数据超时
		return conn, nil
	}

	if qt.proxyip == "" {
		transport = &http.Transport{Dial: Dial, ResponseHeaderTimeout: qt.timesReceive}
	} else {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + qt.proxyip)
		}
		transport = &http.Transport{Proxy: proxy, Dial: Dial, ResponseHeaderTimeout: qt.timesReceive}
	}

	client := &http.Client{Transport: transport, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return errors.New("Location Err") //说明这是一个重定向 ，因为禁止的自动重定向 返回一个重定向错误
	}}
	if qt.setRedirect {
		client = &http.Client{Transport: transport}
	}
	compress = false
	resp, err := client.Do(req) //提交请求
	if err != nil {
		if !strings.Contains(err.Error(), "Location Err") { //是否找到自定义的重定向错误信息
			qt.ret.err = "[winhttp][ client.Do err]" + err.Error()
			return
		}
	}
	defer resp.Body.Close()         //在函数结束时，关闭句柄
	qt.ret.Head = resp.Header       //获取返回的协议头
	for k, v := range qt.ret.Head { //遍历协议头
		if qt.gzip { //是否开启Gip自动解压
			if strings.Contains(k, "encoding") {
				for l := 0; l < len(v); l++ { //从返回的协议头中，去重Cookie进行合并更新
					if strings.Contains(v[l], "qtgzip") {
						compress = true //返回协议头中是否有 gzip字样
						break
					}
				}
			}
		}

		if k == "Set-Cookie" { //如果协议头有设置Cookie的
			for l := 0; l < len(v); l++ { //从返回的协议头中，去重Cookie进行合并更新
				mstr := takeMiddle("Set-Cookie"+v[l], "Set-Cookie", ";")
				arr := strings.Split(mstr, "=")
				if len(arr) == 2 {
					qt.cookies[arr[0]] = arr[1]
				}

			}
		}
	}
	qt.ret.code = resp.StatusCode        //获取返回的状态码
	body, _ := ioutil.ReadAll(resp.Body) //获取返回的Body
	if qt.gzip {                         //是否开启gzip 自动解压
		if compress { //返回协议头中是否有标志出Gzip
			body = NewGzip().Uncompress(body) //解压缩
		} else if len(body) > 8 {
			if body[0] == 31 && body[1] == 139 && body[2] == 8 && body[3] == 0 && body[4] == 0 && body[5] == 0 && body[6] == 0 && body[7] == 0 && body[8] == 0 {
				body = NewGzip().Uncompress(body) //解压缩
			}
		}
	}
	qt.ret.ret = body

}

// takeMiddle .
func takeMiddle(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	str = str[s : s+e]
	return str
}

// New .创建一个新的实例
func newWinhttp() *winhttp {
	tmp := new(winhttp)
	tmp.setRedirect = true                //允许自动重定向
	tmp.timesConnect = 15 * time.Second   //默认连接超时15秒
	tmp.timesSend = 15 * time.Second      //默认发送超时15秒
	tmp.timesReceive = 30 * time.Second   //默认响应时间
	tmp.head = make(map[string]string)
	tmp.head2 = make(map[string]string)
	tmp.cookies = make(map[string]string)
	return tmp
}
