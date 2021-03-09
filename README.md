# qt

> * #  封装有以下功能
>> #### JSON 
>> #### HTTP 
>> #### mysql 
>> #### 常用随机功能 
>> #### 常用字符串操作 
>> #### Gzip 
>> #### Zlib 
>> #### Socket Server 
>> #### Socket Client 
>> #### RSA 算法 
>> #### AES 算法 
>> #### DES 算法 
>> #### Hash加密 

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
