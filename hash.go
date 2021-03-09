package qt

import (
	"crypto"
	"crypto/hmac"
)

func Sha512_256(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}

	if skey != nil {
		hash := hmac.New(crypto.SHA512_256.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA512_256.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha512(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA512.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA512.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha512_224(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA512_224.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA512_224.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha384(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA384.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA384.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha256(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA256.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA256.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha224(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA224.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA224.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Md4(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.MD4.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.MD4.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */

func Md5(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}

	if skey != nil && len(skey) > 0 {
		hash := hmac.New(crypto.MD5.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	}
	hash := crypto.MD5.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}

/*
*======================================================
 */
func Sha1(Data, key interface{}) []byte {
	signByte := []byte{}
	switch bin := Data.(type) {
	case []byte:
		signByte = bin
	case string:
		signByte = []byte(bin)
	default:
		return signByte
	}
	skey := []byte{}
	switch bin := key.(type) {
	case []byte:
		skey = bin
	case string:
		skey = []byte(bin)
	default:
		skey = nil
	}
	if skey != nil {
		hash := hmac.New(crypto.SHA1.New, skey)
		hash.Write(signByte)
		return hash.Sum(nil)
	} else if key != nil {
		return skey
	}
	hash := crypto.SHA1.New()
	hash.Write(signByte)
	return hash.Sum(nil)
}
