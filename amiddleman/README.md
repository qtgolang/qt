> ## HttpAmiddleman
> ## HTTP/HTTPS 中间人抓包
> #### 转自 ：<a href="github.com/nicecp/GoIyov">   https://github.com/nicecp/GoIyov </a> 修改

 ```
package main

import (
	"fmt"
	"github.com/qtgolang/amiddleman"
)
type Handler struct {
}

func (handler *Handler) BeforeRequest(entity *entity.Entity) {
	entity.Request.Header.Set("Accept-Encoding", "") //设置Head

	Mod := entity.Request.Method
	Host := entity.Request.Host
	Path := entity.Request.RequestURI
	buf := new(bytes.Buffer)
	buf.ReadFrom(entity.GetRequestBody())
	Body := buf.String() 
	fmt.Println("请求 Mod", Mod)
	fmt.Println("请求 Host", Host)
	fmt.Println("请求 Path", Path)
	fmt.Println("请求 Body len", len(Body))
	fmt.Println("请求 Body", Body) 

}
func (handler *Handler) BeforeResponse(entity *entity.Entity, err error) {
	Mod := entity.Request.Method
	Host := entity.Request.Host
	Path := entity.Request.RequestURI
	buf := new(bytes.Buffer)
	buf.ReadFrom(entity.GetResponseBody())
	Body := buf.String()

	fmt.Println("Ret Mod", Mod)
	fmt.Println("Ret Host", Host)
	fmt.Println("Ret Path", Path)
	fmt.Println("Ret Body len", len(Body))
	fmt.Println("Ret Body", Body) 
}
func (handler *Handler) ErrorLog(err error) {}

func main() {
	amiddleman.Stat(8080,&Handler{},amiddleman.RootCa,amiddleman.RootKey)
}
 ```
 