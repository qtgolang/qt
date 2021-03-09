package qt

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type rand_ struct {
	Rand *rand.Rand
}

//AndrodidSdkVersion 随机获取 安卓版本信息
func (qtrand *rand_) AndrodidSdkVersion() (安卓版本 string, 安卓级别 int, 版本号 string) {
	arr := strings.Split("8.1.0|8.0.0|7.1.2|7.1.1|7.1.0|7.0.0|6.0.1|6.0.0|5.1.1|5.1.0|5.0.2|5.0.1|5.0.0|4.4.4|4.4.3|4.4.2|4.4.1|4.4|4.3.1|4.3|4.2.2|4.2.1|4.2|4.1.2|4.1.1|4.0.4|4.0.3|4.0.2|4.0.1", "|")
	安卓版本 = arr[qtrand.RandoMnumber(0, len(arr)-1)]
	switch 安卓版本 {

	case "8.1.0":
		arr1 := strings.Split("OPM5.171019.015|OPM3.171019.014|OPM1.171019.019|OPM1.171019.018|OPM1.171019.016|OPM5.171019.014|OPM2.171019.016|OPM3.171019.013|OPM1.171019.015|OPM1.171019.014|OPM1.171019.013|OPM1.171019.012|OPM2.171019.012|OPM1.171019.011", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 27
	case "8.0.0":
		arr1 := strings.Split("OPR5.170623.014|OPR4.170623.020|OPD3.170816.023|OPD1.170816.025|OPR6.170623.023|OPR5.170623.011|OPR3.170623.013|OPR2.170623.027|OPR1.170623.032|OPD3.170816.016|OPD2.170816.015|OPD1.170816.018|OPD3.170816.012|OPD1.170816.012|OPD1.170816.011|OPD1.170816.010|OPR5.170623.007|OPR4.170623.009|OPR3.170623.008|OPR1.170623.027|OPR6.170623.021|OPR6.170623.019|OPR4.170623.006|OPR3.170623.007|OPR1.170623.026|OPR6.170623.013|OPR6.170623.012|OPR6.170623.011|OPR6.170623.010", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 26
	case "7.1.2":
		arr1 := strings.Split("N2G48H|NZH54D|NKG47S|NHG47Q|NJH47F|N2G48C|NZH54B|NKG47M|NJH47D|NHG47O|N2G48B|N2G47Z|NJH47B|NJH34C|NKG47L|NHG47N|N2G47X|N2G47W|NHG47L|N2G47T|N2G47R|N2G47O|NHG47K|N2G47J|N2G47H|N2G47F|N2G47E|N2G47D", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 25
	case "7.1.1":
		arr1 := strings.Split("N9F27M|NGI77B|N6F27M|N4F27P|N9F27L|NGI55D|N4F27O|N8I11B|N9F27H|N6F27I|N4F27K|N9F27F|N6F27H|N4F27I|N9F27C|N6F27E|N4F27E|N6F27C|N4F27B|N6F26Y|NOF27D|N4F26X|N4F26U|N6F26U|NUF26N|NOF27C|NOF27B|N4F26T|NMF27D|NMF26X|NOF26W|NOF26V|N6F26R|NUF26K|N4F26Q|N4F26O|N6F26Q|N4F26M|N4F26J|N4F26I|NMF26V|NMF26U|NMF26R|NMF26Q|NMF26O|NMF26J|NMF26H|NMF26F", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 25
	case "7.1.0":
		arr1 := strings.Split("NDE63X|NDE63V|NDE63U|NDE63P|NDE63L|NDE63H", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 25
	case "7.0.0":
		arr1 := strings.Split("NBD92Q|NBD92N|NBD92G|NBD92F|NBD92E|NBD92D|NBD91Z|NBD91Y|NBD91X|NBD91U|N5D91L|NBD91P|NRD91K|NRD91N|NBD90Z|NBD90X|NBD90W|NRD91D|NRD90U|NRD90T|NRD90S|NRD90R|NRD90M", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 24
	case "6.0.1":
		arr1 := strings.Split("MOI10E|MOB31Z|MOB31T|MOB31S|M4B30Z|MOB31K|MMB31C|M4B30X|MOB31H|MMB30Y|MTC20K|MOB31E|MMB30W|MXC89L|MTC20F|MOB30Y|MOB30X|MOB30W|MMB30S|MMB30R|MXC89K|MTC19Z|MTC19X|MOB30P|MOB30O|MMB30M|MMB30K|MOB30M|MTC19V|MOB30J|MOB30I|MOB30H|MOB30G|MXC89H|MXC89F|MMB30J|MTC19T|M5C14J|MOB30D|MHC19Q|MHC19J|MHC19I|MMB29X|MXC14G|MMB29V|MXB48T|MMB29U|MMB29R|MMB29Q|MMB29T|MMB29S|MMB29P|MMB29O|MXB48K|MXB48J|MMB29M|MMB29K", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 23
	case "6.0.0":
		arr1 := strings.Split("MMB29N|MDB08M|MDB08L|MDB08K|MDB08I|MDA89E|MDA89D|MRA59B|MRA58X|MRA58V|MRA58U|MRA58N|MRA58K", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 23
	case "5.1.1":
		arr1 := strings.Split("LMY49M|LMY49J|LMY49I|LMY49H|LMY49G|LMY49F|LMY48Z|LYZ28N|LMY48Y|LMY48X|LMY48W|LVY48H|LYZ28M|LMY48U|LMY48T|LVY48F|LYZ28K|LMY48P|LMY48N|LMY48M|LVY48E|LYZ28J|LMY48J|LMY48I|LVY48C|LMY48G|LYZ28E|LMY47Z|LMY48B|LMY47X|LMY47V", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 22
	case "5.1.0":
		arr1 := strings.Split("LMY47O|LMY47M|LMY47I|LMY47E|LMY47D", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 22
	case "5.0.2":
		arr1 := strings.Split("LRX22L|LRX22G", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 21
	case "5.0.1":
		arr1 := strings.Split("LRX22C", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 21
	case "5.0.0":
		arr1 := strings.Split("LRX21V|LRX21T|LRX21R|LRX21Q|LRX21P|LRX21O|LRX21M|LRX21L", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 21
	case "4.4.4":
		arr1 := strings.Split("KTU84Q|KTU84P", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 19
	case "4.4.3":
		arr1 := strings.Split("KTU84M|KTU84L", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 19
	case "4.4.2":
		arr1 := strings.Split("KVT49L|KOT49H", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 19
	case "4.4.1":
		arr1 := strings.Split("KOT49E", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 19
	case "4.4":
		arr1 := strings.Split("KRT16S|KRT16M", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 19
	case "4.3.1":
		arr1 := strings.Split("JLS36I", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 18
	case "4.3":
		arr1 := strings.Split("JLS36C|JSS15R|JSS15Q|JSS15J|JSR78D|JWR66Y|JWR66V|JWR66N|JWR66L", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 18
	case "4.2.2":
		arr1 := strings.Split("JDQ39E|JDQ39B|JDQ39", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 17
	case "4.2.1":
		arr1 := strings.Split("JOP40G|JOP40F|JOP40D", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 17
	case "4.2":
		arr1 := strings.Split("JOP40C", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 17
	case "4.1.2":
		arr1 := strings.Split("JZO54M|JZO54L|JZO54K", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 16
	case "4.1.1":
		arr1 := strings.Split("JRO03S|JRO03R|JRO03O|JRO03L|JRO03H|JRO03E|JRO03D|JRO03C", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 16
	case "4.0.4":
		arr1 := strings.Split("IMM76L|IMM76K|IMM76I|IMM76D|IMM76", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 15
	case "4.0.3":
		arr1 := strings.Split("IML77|IML74K", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 15
	case "4.0.2":
		版本号 = "ICL53F"
		安卓级别 = 14
	case "4.0.1":
		arr1 := strings.Split("ITL41F|ITL41D|ITL41D", "|")
		版本号 = arr1[qtrand.RandoMnumber(0, len(arr1)-1)]
		安卓级别 = 14

	}
	return
}

// Char .
//
// 随机多少个字符
func (qtrand *rand_) Char(num int) string {

	if num < 1 {
		return ""
	}
	tmp := []byte{113, 97, 122, 119, 115, 120, 101, 100, 99, 114, 102, 118, 116, 103, 98, 121, 104, 110, 117, 106, 109, 105, 107, 111, 108, 112, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 81, 65, 90, 87, 83, 88, 69, 68, 67, 82, 70, 86, 84, 71, 66, 89, 72, 78, 85, 74, 77, 73, 75, 79, 76, 80}
	return qtrand.srand(num, tmp)
}

// GetTimeUnix 获取当前时间戳 参数若为 true 获取10 位的时间戳，否则为13位时间戳
func (qtrand *rand_) GetTimeUnix(val bool) int64 {
	if val {
		return time.Now().UnixNano() / 1000 / 1000 / 1000
	}
	return time.Now().UnixNano() / 1000 / 1000
}
// TimeStrToTimeUnix 指定时间到时间戳 10位
func (qtrand *rand_) TimeStrToTimeUnix(times string) int64 {

	formatTime,err:=time.Parse("2006-01-02 15:04:05",times)
	if err==nil{
		return formatTime.Unix()
	}
	return -1
}

// UUID .
//
// 随机一个UUID 总长度36
func (qtrand *rand_) UUID() string {
	str := []byte(qtrand.CharHEX(36))
	str[8] = 45
	str[13] = 45
	str[18] = 45
	str[23] = 45
	str[14] = 52 //uuid这一位 必须是4
	return string(str)
}

// CharHEX .
//
// 随机十六进制字符
func (qtrand *rand_) CharHEX(num int) string {
	TMP := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70}
	return strings.ToLower(qtrand.srand(num, TMP))
}

// RandoMnumber .
//
// 取随机返回数字
//
//假设 取 1-5 的范围 那么  出现1 和 5 的情况
func (qtrand *rand_) RandoMnumber(Min, Max int) int {
	return qtrand.Rand.Intn(Max-Min+1) + Min
}

// Number .
//
// 随机多少个数字
func (qtrand *rand_) Number(num int) string {
	if num < 1 {
		return ""
	}
	tmp := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	return qtrand.srand(num, tmp)
}
func (qtrand *rand_) srand(num int, tmp []byte) string {
	if num < 1 {
		return ""
	}
	str := ""
	stmp := []byte{1}
	for i := 0; i < num; i++ {
		stmp[0] = tmp[qtrand.Rand.Intn(len(tmp))]
		str += string(stmp)
	}
	return str
}
func handle_(a string) int64 {
	tmp := Md5(a, nil)
	if len(tmp) != 16 {
		return 0
	}
	return abs(int64(binary.BigEndian.Uint64(tmp)))
}


//初始化一个随机类
func newRand_() *rand_ {
	tmp := new(rand_)
	var i int64
	i = handle_(fmt.Sprintf("%d", &i) + fmt.Sprintf("%d", &tmp) + strconv.FormatInt(time.Now().UnixNano(), 10))
	s1 := rand.NewSource(i)
	tmp.Rand = rand.New(s1)
	return tmp
}

// MAC .
//
//随机一个MAC
func (qtrand *rand_) MAC() string {
	return qtrand.CharHEX(2) + ":" + qtrand.CharHEX(2) + ":" + qtrand.CharHEX(2) + ":" + qtrand.CharHEX(2) + ":" + qtrand.CharHEX(2) + ":" + qtrand.CharHEX(2)
}

// IMEI .
//
//随机一个IMEI
func (qtrand *rand_) IMEI() string {
	_ReportingBodyIdentifier := []string{"01", "35", "86"} //可能还有更多，现在默认3个就行了
	_TAC2 := qtrand.Number(4)
	_FAC := qtrand.Number(2)
	_IMEI := _ReportingBodyIdentifier[qtrand.RandoMnumber(0, len(_ReportingBodyIdentifier)-1)] + _TAC2 + _FAC + qtrand.Number(6)
	return _IMEI + strconv.Itoa(luhn_计算IMEI校验码(_IMEI))
}
func luhn_计算IMEI校验码(s string) int {
	_len := len(s)
	ji := 0
	he := 0
	_he := 0
	base := 10
	_int := 0
	for i := 1; i <= _len; i++ {
		_int, _ = strconv.Atoi(s[i-1 : i])
		if i%2 == 0 {
			_int = _int * 2
			he = he + iDiv(_int, base) + _int%base
		} else {
			ji = ji + _int
		}

	}
	_he = he + ji
	yu := _he % base
	if yu == 0 {
		return 0
	}
	return int((iDiv(_he, base)+1)*base - _he)
}

// IDiv .
//
// 整除
func iDiv(a, b int) int {
	c := a % b
	d := a - c
	i := 0
	for x := 0; x < a; x++ {
		d = d - b
		i++
		if d == 0 {
			return i
		}
	}
	return 0
}
