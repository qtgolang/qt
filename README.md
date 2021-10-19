# qt

> * #  封装有以下功能
>> #### <a href="https://github.com/qtgolang/qt/blob/master/Json.go">JSON</a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qthttp.go">HTTP </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtmysql.go">mysql  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtrand.go">常用随机功能  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtString.go">常用字符串操作  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/gzip.go">Gzip  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtzlib.go">Zlib  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtsocket.go">Socket Server  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtsocket.go">Socket Client  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtrsa.go">RSA 算法  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtaes.go">AES 算法  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtdes.go">DES 算法  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/hash.go">Hash加密  </a>
>> #### <a href="https://github.com/qtgolang/qt/blob/master/qtfile.go">常用文件操作  </a>

 * # JSON 示例
 ```
package main

import (
	"fmt"
	"github.com/qtgolang/qt"
)

func main() {
	j:=qt.JSON()
	j.Untie(`{"a":1,"b":true}`)
	fmt.Println(j.GetFloat64("a"))
	fmt.Println(j.GetBool("b"))
    }
 ```

 * # Aes 示例 /Des 类似
 ```
package main

import (
	"fmt"
	"github.com/qtgolang/qt"
)

func main() {
	a := qt.Aes()
	a.SetPadding(qt.Type_Const_Padding_Pkcs5)
	a.SetEncMethod(qt.Type_Const_AES_DES_ECB)
	a.SetFill(true)
	a.SetKey("01234567890123456")
	fmt.Println(hex.EncodeToString(a.Encrypt([]byte("000000000"))))
    }
 ```

 * # Http 示例 
 ```
package main

import (
	"fmt"
	"github.com/qtgolang/qt"
)

func main() {
	h := qt.Newhttp()
	h.Open("GET", "https://www.baidu.com")
	h.Send("")
	fmt.Println(h.GetBodyAuto())
    }
 ```
