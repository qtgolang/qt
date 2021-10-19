package qt

/*
#include <stdlib.h>
*/
import "C"

import (
	"syscall"
	"unsafe"
)

// Call
//
// Call 一个内存地址 只能接收 uintptr int int8 int16 int32 int64  string []byte bool 类型
// 返回-1 表示失败 返回其他并表示执行成功  返回 Call 的目标函数的返回值返回数据的指针
func Call(address int, arg ...interface{}) uintptr {

	var args []uintptr
	var Frees []*C.char
	for _, name := range arg {
		switch val := name.(type) {
		case uintptr:
			args = append(args, val)
		case int:
			args = append(args, uintptr(val))
		case int8:
			args = append(args, uintptr(val))
		case int16:
			args = append(args, uintptr(val))
		case int32:
			args = append(args, uintptr(val))
		case int64:
			args = append(args, uintptr(val))
		case bool:
			if val {
				args = append(args, uintptr(1))
			} else {
				args = append(args, uintptr(0))
			}
		case string:
			n := C.CString(val)
			Frees = append(Frees, n)
			args = append(args, uintptr(unsafe.Pointer(n)))
		case []byte:
			n := C.CString(string(val))
			Frees = append(Frees, n)
			args = append(args, uintptr(unsafe.Pointer(n)))
		default:
			return -1 //如果有其他参数类型 直接报错返回
		}
	}
	slen := len(args)
	for index := 0; index < (18 - slen); index++ {
		args = append(args, uintptr(0))
	}
	r1, _, _ := syscall.Syscall18(uintptr(address), uintptr(slen), args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17])
	for index := 0; index < len(Frees); index++ {
		C.free(unsafe.Pointer(Frees[index]))
	}
	return r1
}