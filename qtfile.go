package qt

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type QtFile struct {
	FileNumer  *os.File
	FileReader *bufio.Reader
}

//CheckFileIsExist
//
// 判断文件是否存在  存在返回 true 不存在返回false
func (c QtFile) CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// Open
//
// 打开一个文件 成功返回 nil
//
// Filename 文件名
//
// OpenType 打开类型 1 清空后打开 2 = 如果已经存在，则在尾部添加写 3 = 如果已经存在，会覆盖写，不会清空原来的文件，而是从头直接覆盖写 4=如果已经存在，则失败
//
// FileMode 其他进程权限 1=无限制 2=禁止读 3=禁止写 4=禁止读写
func (c *QtFile) Open(Filename string, OpenType, FileMode int) error {
	_OpenType := OpenType
	_permissions := os.FileMode(0777)
	if _OpenType == 1 {
		_OpenType = os.O_CREATE | os.O_APPEND | os.O_RDWR
		c.RemoveFile(Filename)
	} else if _OpenType == 2 {
		_OpenType = os.O_CREATE | os.O_APPEND | os.O_RDWR
	} else if _OpenType == 3 {
		_OpenType = os.O_CREATE
	} else if _OpenType == 4 {
		_OpenType = os.O_CREATE | os.O_EXCL | os.O_RDWR
	}
	if FileMode == 1 {
		_permissions = 0777
	} else if FileMode == 2 {
		_permissions = 0772
	} else if FileMode == 3 {
		_permissions = 0774
	} else {
		_permissions = 0770
	}
	var err error
	c.FileNumer, err = os.OpenFile(Filename, _OpenType, _permissions)
	c.FileReader = bufio.NewReader(c.FileNumer)
	return err
}

// Close
//
// 关闭文件
func (c *QtFile) Close() error {
	return c.FileNumer.Close()
}

// Write
// 写[]byte
func (c *QtFile) Write(body []byte) error {
	_, err := c.FileNumer.Write(body)
	return err
}

// WriteString
//
// 写入Sting
func (c *QtFile) WriteString(Str string) error {
	_, err := c.FileNumer.WriteString(Str)
	return err
}

// WriteStringLine
// 写一行Sting 尾部自动带\r\n
func (c *QtFile) WriteStringLine(Str string) error {
	return c.Write([]byte(Str + "\r\n"))
}

// ReadString
//
// 读一行字符串
func (c *QtFile) ReadString() (string, error) {
	line, err := c.FileReader.ReadString('\n')
	if len(line) < 1 {
		return line, err
	}
	if line[len(line)-1:] == "\n" {
		line = line[:len(line)-1]
	}
	if len(line) < 1 {
		return line, err
	}
	if line[len(line)-1:] == "\r" {
		line = line[:len(line)-1]
	}
	return line, err
}

// WriteStringtoFile
//
// 写字符串到文件
func (c QtFile) WriteStringtoFile(Text, Filename string) error {
	return c.WriteBytestoFile([]byte(Text), Filename)
}

// RemoveFile
//
// 删除文件
func (c QtFile) RemoveFile(Filename string) error {
	return os.Remove(Filename)
}

// WriteBytestoFile
//
// 写[]byte到文件
func (c QtFile) WriteBytestoFile(bytes []byte, Filename string) error {
	var f *os.File
	var err error
	//文件是否存在
	if c.CheckFileIsExist(Filename) {
		//存在 删除
		err = c.RemoveFile(Filename)
		if err != nil {
			return err
		}
	}
	//创建文件
	f, err = os.Create(Filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// 写入
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

// ReadFileString
//
// 读入文件并转为String
func (c QtFile) ReadFileString(Filename string) string {
	name := strings.ReplaceAll(Filename, "\"", "")
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return ""
	}
	return string(b)
}

// ReadFilebytes
//
// 读入文件返回字节数组
func (c QtFile) ReadFilebytes(Filename string) []byte {
	name := strings.ReplaceAll(Filename, "\"", "")
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte("")
	}
	return (b)
}

// CreateDirAll
//
// 创建目录 例如d:\1\2\3\4 只需要 D 盘存在即可
func (c QtFile) CreateDirAll(Filename string) error {
	return os.MkdirAll(Filename, 0777)
}

// CreateDir
//
// 创建目录 例如d:\1\2\3\4   目录 d:\1\2\3\ 必须存在
func (c QtFile) CreateDir(Filename string) error {
	return os.Mkdir(Filename, 0777)
}
