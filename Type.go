package qt

const (
	// Hash_Stat Hash开始处的标志
	Type_Const_Hash_Stat       = 1
	Type_Const_Hash_Sha1       = 1
	Type_Const_Hash_Md4        = 2
	Type_Const_Hash_Md5        = 3
	Type_Const_Hash_Sha224     = 4
	Type_Const_Hash_Sha256     = 5
	Type_Const_Hash_Sha384     = 6
	Type_Const_Hash_Sha512     = 7
	Type_Const_Hash_Sha512_224 = 8
	Type_Const_Hash_Sha512_256 = 9
	// Hash_End Hash结束处的标志
	Type_Const_Hash_End = 9

	Type_Const_Padding_Pkcs5     = 0
	Type_Const_Padding_Pkcs7     = 1
	Type_Const_Padding_NoPadding = 2
	Type_Const_Padding_Zero      = 3
	Type_Const_Padding_Iso10126  = 4
	Type_Const_Padding_Iso97971  = 5
	Type_Const_Padding_AnsiX923  = 6
	Type_Const_AES_DES_ECB       = 0
	Type_Const_AES_DES_CBC       = 1
	Type_Const_AES_DES_FCB       = 2
	Type_Const_AES_keySize_128   = 128
	Type_Const_AES_keySize_192   = 192
	Type_Const_AES_keySize_256   = 256
	Type_Const_AES_keySize_Auto  = 0
	Type_Const_DES_3DES_ECB      = 3
	Type_Const_DES_3DES_CBC      = 4
	Type_Const_DES_3DES_FCB      = 5
)
