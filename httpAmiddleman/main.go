package main

import (
	"amiddleman/GoIyov"
	"amiddleman/GoIyov/entity"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

//package httpAmiddleman
type Handler struct {
}

func (handler *Handler) BeforeRequest(entity *entity.Entity) {
	entity.Request.Header.Set("Accept-Encoding", "")

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

	//qt.Call(callback, Mod, Host, Path, Body, len(Body))

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
	//go qt.Call(callback, Mod, Host, Path, Body, len(Body))
}
func (handler *Handler) ErrorLog(err error) {}

var _rootCa = `-----BEGIN CERTIFICATE-----
MIIDizCCAnOgAwIBAgIUKZuAsiiXCMz613rrURfxAHNuU7YwDQYJKoZIhvcNAQEL
BQAwVDELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAkJKMQswCQYDVQQHDAJCSjENMAsG
A1UECgwEbGl2ZTENMAsGA1UECwwEUk9PVDENMAsGA1UEAwwEUk9PVDAgFw0yMTEw
MjEwMzUwNTBaGA8yMTIxMDkyNzAzNTA1MFowVDELMAkGA1UEBhMCQ04xCzAJBgNV
BAgMAkJKMQswCQYDVQQHDAJCSjENMAsGA1UECgwEbGl2ZTENMAsGA1UECwwEUk9P
VDENMAsGA1UEAwwEUk9PVDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AMI+CTQSiFlew0i5Ir0OczDnucnHbCoCEb5m3x3VyH+KFAPdc8DltkETlSlvKbcc
C3G4Tm8IQ0lN9GpFepqbd/+06vhiZg+NU10hoXFP7pk3LwnrBXFHSHAN90y/7REP
1wPRdDUglI6WJKvEailZXdNgch+HTNXDlbpAKytqRe2m/pdzqTDWpiPu1UFQ0AtQ
rRhk+rCQe467cW7ekY/t7pEXqEwxUqKjHWvCXUAFUy6rOYe6/VZJWQF7Wo4zyidy
SYCF2t7RSpMXiwt5oEnuMT+L6g9RG+5DOpHJmg5LsqCYaQ+UWB2sthIlUni7D5E4
DFlKi+1AzV3sn6qrzAE0wlMCAwEAAaNTMFEwHQYDVR0OBBYEFL6bnl/tCNaISSbE
IZEkxBZr0kZEMB8GA1UdIwQYMBaAFL6bnl/tCNaISSbEIZEkxBZr0kZEMA8GA1Ud
EwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAIIeZjAJ1opJQnNPbarbYFTG
I1yx3IbBjwr0/TX6Ku+BFU2McLSQFYdE8mtzcKGUryz15fBnirr+Xtjp7oy+1sZp
fiJToAw0fob/hjOWrLd5Oa8ZbwCCIbqauQy75SD4T/zBGAnyUSvTvhlITKPDrvMg
q/78o+3hitOCzytPgy9njee45DyhzW3/kczTXiBObyUrgI2YnFsbBBunYu5r5atb
jflL/BPXbo2I6AUSIZBYP+sjlDugttxtm7dao61fiMREkd5sgFhFi1b7HAETESKs
YuUTLN7fP7MUIOy+1cf+uX8STUTUz5tmu4eH6NiG4+HVdQHWQ0fI/3RoZmvghOg=
-----END CERTIFICATE-----
`
var _rootKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwj4JNBKIWV7DSLkivQ5zMOe5ycdsKgIRvmbfHdXIf4oUA91z
wOW2QROVKW8ptxwLcbhObwhDSU30akV6mpt3/7Tq+GJmD41TXSGhcU/umTcvCesF
cUdIcA33TL/tEQ/XA9F0NSCUjpYkq8RqKVld02ByH4dM1cOVukArK2pF7ab+l3Op
MNamI+7VQVDQC1CtGGT6sJB7jrtxbt6Rj+3ukReoTDFSoqMda8JdQAVTLqs5h7r9
VklZAXtajjPKJ3JJgIXa3tFKkxeLC3mgSe4xP4vqD1Eb7kM6kcmaDkuyoJhpD5RY
Hay2EiVSeLsPkTgMWUqL7UDNXeyfqqvMATTCUwIDAQABAoIBADAr7GJimjf3vVyC
tW/Hbp/ZSAUXlOaSHs5merzOcOvYVYBj4jobfeiDr8FX1TOZG+l2+NSmavF6CLx2
QKKpHchnmMJCvObseJknJv+QAC81NVAzXeUq9Xb8r7D0BBnwILXMsxId41m9OJj0
nBNnG1K5n3PcPRziZXaMRWh98ZuvGE2efXHO+Z+bo1O3wyBb7aH7dwMEXLRwUxMO
lRrcczLdySejcuc+z0p9jJUhCnXAuuunOhjVrv/jVsuqUaFfiuIvp0A2dLXu/Bea
mWywSr+WqTZdszErghXuEjmBMK7rFz84/fvqxUMYM+rI2PWhjIOw0K0sSkIK0SwR
wrVdFWkCgYEA9311dONQfJiwxNGpSqJaI+XdNkuzyCFollxxbJVLBrnBd5+rt3bX
4wJwttQrIDY9mlGtitijf43C5yitO9Y9z4asON8hgjfalBp2/g98Aym0oLr+GHaJ
OQhVbpHawYkrKQQUXrT2UNrUI88LJM4Ezxs1+pvB01cMVc27DxqunF0CgYEAyOvc
m0cbWS5rBLY7oTFYebMqreDY/PCXbWBeDXCuTvEa/24HRmuK3IPsv5ircOgHSxBH
imZXBMITev0MPseiMYajqcgnSUWNzEtczrBMi00ezIC2OExifWExIaYAWeF7vQoY
JAxaVAy4IVZ2gkdr2Ib4ckoH1GyUKS5/8uybbm8CgYEA08IiHOhetq1DGrS2IGj3
rZ2BgHKXmlaTkYv5dMnszw1jb0JMgAcMw20UGitB1ybx4LegQJwKkRovuPACAZ/X
dViqxWKN3kiCGpTmWY3QMzidF6XHwhCyav0pdBkSTuRZ7JdUApDd7OK//v+pbk1k
qfyDXDGnT3g80rHcKxlOa/UCgYAy8mvMC+nJYVXNqg/Qvdx7b40A7iTXboJXP7pZ
mhr49XYtEs9Rce+SHvmnU6UVSiCfTg917qFeGYArDYNPX/ump0dUw+YCVFqyVOHB
nz6pi/KlVHXgLK5EsKJur9Mi92QS+T5J1cAJ0/fUrEH8ovObwg25nUADA01Ga/4R
sSFwjQKBgQDcGD7f7dMvt5azi7/OkAc18Z7NxWwyZ52i/JwtejOoO1ZhY1R0zT1U
hu3xdcZQ98EHSzY9i/cnj8hbpRc158NRTIJ9+16XFtw4mKXwnEetw4cnotF8rUAs
EkvK5BlNCpvFCO8TKImcB/g4v8R7FNxmM4tUE1HJzSym4jbiaWtf8w==
-----END RSA PRIVATE KEY-----
`

func main() {
	proxy := GoIyov.NewWithDelegate(&Handler{}, _rootCa, _rootKey)
	server := &http.Server{
		Addr: ":8880",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			proxy.ServerHandler(rw, req)
		}),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
