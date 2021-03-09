package qt

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type qtGzip struct {
	Err string
}

func NewGzip() *qtGzip {
	tmp := new(qtGzip)
	return tmp
}
func (qt *qtGzip) Compress(in []byte) []byte {
	qt.Err = ""
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		qt.Err = err.Error()
		return out
	}
	err = writer.Close()
	if err != nil {
		qt.Err = err.Error()
		return out
	}

	return buffer.Bytes()
}

func (qt *qtGzip) Uncompress(in []byte) []byte {
	qt.Err = ""
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		qt.Err = err.Error()
		return out
	}
	defer reader.Close()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		qt.Err = err.Error()
		var out []byte
		return out
	}
	return data
}
