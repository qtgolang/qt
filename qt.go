package qt

//String 返回一个Qtstring类型
func String() *qtstring {
	return newstring()
}

//Winhttp 返回一个Winhttp类型
//
//实现GET POST 等请求
func Http() *winhttp {
	return newWinhttp()
}

// Gzip 返回一个Gzip类型
//
//可以用于 压缩和解压缩
func Gzip() *qtGzip {
	return NewGzip()
}

// Hash 常用Hash 算法
//
//参数1 请使用 [Type.Hash_]选择
func Hash(HashType int, data interface{}) []byte {
	return Hmac(HashType, data, nil)
}

// Hmac 常用Hash-Hmac 算法
//
//参数1 请使用 [Type_Const_Hash_]选择
func Hmac(HashType int, data, key interface{}) []byte {
	if Type_Const_Hash_Stat < 1 || Type_Const_Hash_End > 9 {
		return []byte{}
	}
	switch HashType {
	case Type_Const_Hash_Sha1:
		return Sha1(data, key)
	case Type_Const_Hash_Md4:
		return Md4(data, key)
	case Type_Const_Hash_Md5:
		return Md5(data, key)
	case Type_Const_Hash_Sha224:
		return Sha224(data, key)
	case Type_Const_Hash_Sha256:
		return Sha256(data, key)
	case Type_Const_Hash_Sha384:
		return Sha384(data, key)
	case Type_Const_Hash_Sha512:
		return Sha512(data, key)
	case Type_Const_Hash_Sha512_224:
		return Sha512_224(data, key)
	case Type_Const_Hash_Sha512_256:
		return Sha512_256(data, key)
	}
	return []byte{}
}

// Json 返回一个Json类型
//
//支持以下类型
//
//动态添加 【map对象/任意类型合集/数组对象/浮点数/整数/逻辑/字符串】 包括在数组内指定位置添加，能自动扩容
//
//动态删除 【map对象/任意类型合集/数组对象/浮点数/整数/逻辑/字符串】 包括在数组内指定位置删除
//
//动态取值 【map对象/任意类型合集/数组对象/浮点数/整数/逻辑/字符串】
func JSON() *qtJson {
	a := newJson()
	a.Untie("{}")
	return a
}

// Aes 返回一个AES类型
func Aes() *qtAes {
	return NewAes()
}

// Des 返回一个DES类型
func Des() *qtDes {
	return NewDes()
}

// Mysql 返回一个Mysql类型
func Mysql() *mysql {
	return newMysql()
}

// SocketServer .
//
//返回一个Socket 的服务端
//
//var Connect = func(server *qtsocket.Server, handle int64) {   }
//
//var Exit = func(server *qtsocket.Server, handle int64, tmpdata *qtJson.Json) {
//
////此时,用户的临时数据已经被清除，这里tmpdata,是最后一次反馈
//
//   }
//
//var Arrival = func(server *qtsocket.Server,handle int64, data []byte) {   }
//
//在 Exit 函数中有一个   tmpdata 记录着 用户所有的临时数据
func SocketServer(Connect func(server *Server, handle int64), Exit func(server *Server, handle int64, tmpdata *qtJson), Arrival func(server *Server, handle int64, data []byte)) *Server {
	tmp := &Server{Connect: Connect, Exit: Exit, Arrival: Arrival}
	tmp.SetBit(4096)
	return tmp
}

// SocketClient .
//
//返回一个Socket 的客户端
func SocketClient() *Client {
	tmp := new(Client)
	return tmp
}

// Zlib .
//
//返回一个Zlib 的实例
//
//现在的Zlib不能调整级别后期修改
func Zlib() *qtzlib {
	return newzlib_()
}

// Rand .
//
//返回一个随机类
func Rand() *rand_ {
	tmp := newRand_()
	return tmp
}

// RSA .
//
// New一个RSA实例
func RSA() *Qtrsa {
	tmp := new(Qtrsa)
	return tmp
}

// File .
//
// New一个File实例
func File() *QtFile {
	tmp := new(QtFile)
	return tmp
}
