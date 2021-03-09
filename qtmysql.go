package qt

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"sync"
	"time"
)

type mysql struct {
	ip       string
	prot     int
	user     string
	pass     string
	database string
	err      string
	sql      *sql.DB
	row      []interface{}
	rowid    int
	myLook sync.Mutex
}

func newMysql() *mysql {
	tmp := new(mysql)
	return tmp
}

// SetInfo 设置配置信息
func (qtmysql *mysql) SetInfo(IP string, Prot int, User, Pass, kuming string) {
	qtmysql.user = User
	qtmysql.pass = Pass
	qtmysql.prot = Prot
	qtmysql.ip = IP
	qtmysql.database = kuming

}

// GetErr 获取错误信息
func (qtmysql *mysql) GetErr() string {
	return qtmysql.err
}

// SetConnMaxtime .
//
//最大连接周期，超过时间的连接就关闭连接
//
//单位 分钟
//
//应当在Open之后
//
//如果小于等于0，则一直保持连接
func (qtmysql *mysql) SetConnMaxtime(Times time.Duration) {
	qtmysql.sql.SetConnMaxLifetime(Times * time.Minute)
}

// SetMaxCoons .
//
//设置最大连接数
//
//应当在Open之后
func (qtmysql *mysql) SetMaxCoons(coons int) {
	qtmysql.sql.SetMaxOpenConns(coons)
}

// SetMaxCoons .
//
//设置闲置连接数
//如果  <=0，则不保留空闲连接。
//
//应当在Open之后
func (qtmysql *mysql) SetIdleCoons(coons int) {
	qtmysql.sql.SetMaxIdleConns(coons)
}

// Open .
//
// 建立mysql连接
func (qtmysql *mysql) Open() bool {
	var err error
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", qtmysql.user, qtmysql.pass, "tcp", qtmysql.ip, qtmysql.prot, qtmysql.database)
	//fmt.Println(dsn)
	qtmysql.sql, err = sql.Open("mysql", dsn)
	if err != nil {
		qtmysql.err = fmt.Sprintf("Open mysql failed,err:%v", err)
		return false
	}
	rows, err := qtmysql.sql.Query("show databases")
	if err != nil {
		qtmysql.err = err.Error()
		return false
	} else {
		defer rows.Close()
	}

	return rows != nil
}

// Close .
//
// 断开mysql连接
func (qtmysql *mysql) Close() {
	qtmysql.sql.Close()
}

// GetQueryNum .
//
// 获取Query语句得到的数量
func (qtmysql *mysql) GetQueryNum() int {
	return len(qtmysql.row)
}

// SetRowIdBack .
//
// 设置返回结果索引后退一
func (qtmysql *mysql) SetRowIdBack() {
	qtmysql.rowid--
	if qtmysql.rowid < 0 {
		qtmysql.rowid = 0
	}
}

// SetRowIdPlus .
//
// 设置返回结果索引加一
func (qtmysql *mysql) SetRowIdPlus() {
	qtmysql.rowid++
	if qtmysql.rowid > len(qtmysql.row) {
		qtmysql.rowid--
	}
}

type Result struct {
	result map[string]interface{}
}

// Getint .
//
//从返回结果中获取整数型数据 失败返回0
func (qtmysql *Result) Getint(Name string) int64 {
	_int64, err := strconv.ParseInt(qtmysql.GetString(Name), 10, 64)
	if err != nil {
		return 0
	}
	return _int64
}

// Getfloat .
//
//从返回结果中获取 浮点型 数据 失败返回0
func (qtmysql *Result) Getfloat(Name string) float64 {
	float, err := strconv.ParseFloat(qtmysql.GetString(Name), 64)
	if err != nil {
		return 0
	}
	return float
}

// GetBool .
//
//从返回结果中获取 逻辑型 数据 失败返回 false
//
//只有值为1的时候返回 true 不管是int还是string 只要是为1 就是 true
func (qtmysql *Result) GetBool(Name string) bool {
	float := qtmysql.Getfloat(Name)
	return float == 1 //1为 true 0为false
}

// GetMap .
//
// 获取这一条的全部数据 这是一个map
func (qtmysql *Result) GetMap() map[string]interface{} {
	return qtmysql.result
}

// GetString .
//
//从返回结果中获取 字符串 数据 失败返回 空文本
func (qtmysql *Result) GetString(Name string) string {
	cc := qtmysql.result[Name]
	switch cv := cc.(type) {
	case string:
		return cv
	case []byte:
		return string(cv)
	case int:
		return strconv.Itoa(cv)
	case int8:
		return strconv.Itoa(int(cv))
	case int16:
		return strconv.Itoa(int(cv))
	case int32:
		return strconv.Itoa(int(cv))
	case int64:
		return strconv.Itoa(int(cv))
	case float32:
		return strconv.FormatFloat(float64(cv), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(cv, 'E', -1, 32)
	default:
		fmt.Println(fmt.Sprintf("%T", cv))
	}
	return ""
}

// Getbytes .
//
//从返回结果中获取 字节流 数据 失败返回 nil
func (qtmysql *Result) Getbytes(Name string) []byte {
	cc := qtmysql.result[Name]
	switch cv := cc.(type) {
	case []byte:
		return cv
	default:
		fmt.Println(fmt.Sprintf("%T", cv))
	}
	return nil
}

// GetResult .
//
// 获取返回的结果，如果在此之前调用 Query 那么之前的结果将清空
//
//请使用 SetRowIdBack  SetRowIdPlus  前进后退
func (qtmysql *mysql) GetResult() *Result {
	if qtmysql.rowid >= len(qtmysql.row) {
		return nil
	}
	switch cc := qtmysql.row[qtmysql.rowid].(type) {
	case map[string]interface{}:
		tmp := new(Result)
		tmp.result = cc
		return tmp
	}
	return nil
}

// Exec .
//
//执行mysql语句 非查询
//
//成功返回，受影响的行数
//
//失败返回 -1
func (qtmysql *mysql) Exec(sql string) int64 {
	qtmysql.myLook.Lock()
	defer  qtmysql.myLook.Unlock()
	Conn, err := qtmysql.sql.Begin() //开始事务
	if err != nil {
		qtmysql.err = err.Error()
		return -1
	}
	rows, err := qtmysql.sql.Exec(sql)
	if err != nil {
		Conn.Rollback() //滚回 事务
		qtmysql.err = err.Error()
		return -1
	}
	h, err := rows.RowsAffected()
	Conn.Commit() //事务 提交
	return h
}

// Query .
//
//查询mysql语句,最好不要用本命令来执行 增 删 改 操作
func (qtmysql *mysql) Query(sql string)( bool,error ){
	qtmysql.myLook.Lock()
	defer  qtmysql.myLook.Unlock()
	rows, err := qtmysql.sql.Query(sql)
	qtmysql.row = append([]interface{}{})
	qtmysql.rowid = 0
	if err != nil {
		qtmysql.err = err.Error()
		return false,err
	}
	defer rows.Close()
	for rows.Next() {
		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = rows.Scan(scanArgs...)
		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				switch f := col.(type) {
				case int:
					record[columns[i]] = f
				case int8:
					record[columns[i]] = f
				case int16:
					record[columns[i]] = f
				case int32:
					record[columns[i]] = f
				case int64:
					record[columns[i]] = f
				case float32:
					record[columns[i]] = f
				case float64:
					record[columns[i]] = f
				case bool:
					record[columns[i]] = f
				case []byte:
					record[columns[i]] = f
				default:
					record[columns[i]] = col.([]byte)
				}

			}
		}
		qtmysql.row = append(qtmysql.row, nil)
		qtmysql.row[len(qtmysql.row)-1] = record
	}
	return len(qtmysql.row) > 0,nil
}
