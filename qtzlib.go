package qt

import (
	"bytes"
	"compress/zlib"
	"io"
)

// Qtzlib .
//
// 自定义的zlib数据
type qtzlib struct {
}

// Compress .
//
//压缩数据
func (q *qtzlib) Compress(input []byte) []byte {
	var buf bytes.Buffer
	compressor, err := zlib.NewWriterLevel(&buf, zlib.DefaultCompression)
	if err != nil {
		//catlog.Debug("压缩失败") 这是我自己写的一个log 注释掉
		return input
	}
	compressor.Write(input)
	compressor.Close()
	return buf.Bytes()
}

// UnCompress .
//
// 解压缩数据
func (q *qtzlib) UnCompress(Src []byte) []byte {
	b := bytes.NewReader(Src)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

// New .
//
// New一个实例
func newzlib_() *qtzlib {
	tmp := new(qtzlib)
	return tmp
}
