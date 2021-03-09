package qt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type qtJson struct {
	json    map[string]interface{}
	err     string //运行过程中的错误
	format  bool
	is_type bool
}

//获取运行过程中捕获到的错误
func (qt_json *qtJson) GetErr() string {
	return qt_json.err
}

//是否需要格式化后 To_String()
func (qt_json *qtJson) SetFormat(Format bool) {
	qt_json.format = Format
}

//是否需要打印输出 未知类型
//
//调试使用
func (qt_json *qtJson) IsType(is bool) {
	qt_json.is_type = is
}

//取文本string
func (qt_json *qtJson) GetString(Name string) string {
	val := qt_json.get_All_type(Name)
	switch vv := val.(type) {
	case string: //是否是字符串类型
		return vv
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "GetString")
		}
	}
	return ""
}

// GetByteArr .
//
// 获取字节数组
func (qt_json *qtJson) GetByteArr(Name string) []byte {
	val := qt_json.get_All_type(Name)
	var tmp bytes.Buffer
	switch vv := val.(type) {
	case []uint8:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(vv[i])
		}
		return tmp.Bytes()
	case []uint16:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []uint32:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []uint64:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []int8:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []int16:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []int32:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []int64:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case []int:
		for i := 0; i < len(vv); i++ {
			tmp.WriteByte(byte(vv[i]))
		}
		return tmp.Bytes()
	case string:
		if vv != "" {
			decodeBytes, err := base64.StdEncoding.DecodeString(vv)
			if err == nil {
				return decodeBytes
			}
		}
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "GetByteArr")
		}
	}
	return []byte{}
}

//取 map[string]interface{} 对象
//
//Map中 可能存在字符串，逻辑，浮点数，整数，类型需要自己通过 switch 判断
func (qt_json *qtJson) GetMapObject(Name string) map[string]interface{} {
	val := qt_json.get_All_type(Name)
	switch vv := val.(type) {
	case map[string]interface{}: //判断类型
		return vv
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "GetMapObject")
		}
	}
	return nil
}

// GetMapNum .
//
// 获取成员数量
func (qt_json *qtJson) GetMapNum(Name string) int {
	val := qt_json.get_All_type(Name)
	switch vv := val.(type) {
	case []interface{}: //判断类型
		return len(vv)
	case []string: //判断类型
		return len(vv)
	case []int: //判断类型
		return len(vv)
	case []int8: //判断类型
		return len(vv)
	case []int16: //判断类型
		return len(vv)
	case []int32: //判断类型
		return len(vv)
	case []int64: //判断类型
		return len(vv)
	case []float32: //判断类型
		return len(vv)
	case []float64: //判断类型
		return len(vv)
	case []bool: //判断类型
		return len(vv)
	case []map[string]interface{}:
		return len(vv)
	case []map[int]interface{}:
		return len(vv)
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "GetMapNum")
		}
	}
	return 0
}

//取  interface {} 对象
//
//数组中可能存在字符串，逻辑，浮点数，整数，类型需要自己通过 switch 判断
func (qt_json *qtJson) GetArrObject(Name string) interface{} {
	val := qt_json.get_All_type(Name)
	return val
}

//取所有 float\Iin 类型 的值 全部转为 float64 失败返回0
func (qt_json *qtJson) GetFloat64(Name string) float64 {
	val := qt_json.get_All_type(Name)
	switch vv := val.(type) {
	case float64: //是否是float64
		return (vv)
	case float32: //是否是float32
		return float64(vv)
	case int: //是否是int
		return float64(vv)
	case int8: //是否是int8
		return float64(vv)
	case int16: //是否是int16
		return float64(vv)
	case int32: //是否是int32
		return float64(vv)
	case int64: //是否是int64
		return float64(vv)
	case string: //是否是int64
		L,_:=strconv.Atoi(vv)
		return float64(L)
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "Get_float64")
		}
	}
	return 0
}

//取逻辑值
//
//若不存在返回 false
func (qt_json *qtJson) GetBool(Name string) bool {
	val := qt_json.get_All_type(Name)
	switch v := val.(type) {
	case bool: //是否是逻辑值
		return v
	}
	return false
}

//设置任意值
//
//如果为 []byte   自动转为Base64 字符串
func (qt_json *qtJson) SetVal(Name string, val interface{}) {
	i := val
	switch vv := val.(type) {
	case int: //是否是int
		i = float64(vv)
	case int8: //是否是int8
		i = float64(vv)
	case int16: //是否是int16
		i = float64(vv)
	case int32: //是否是int32
		i = float64(vv)
	case int64: //是否是int64
		i = float64(vv)
	}
	qt_json.set_all_type(Name, i)
}

// 暂时将 \+  换成.
func txtReplace(Name string) string {
	return strings.Replace(strings.Replace(strings.Replace(Name, "`-/", "[", -1), "`+/", "]", -1), "\\+", ".", -1)
}

//删除字段，或字段值
func (qt_json *qtJson) Delete(Name string) {
	Name = strings.Replace(Name, "\\.", "\\+", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	Name = strings.Replace(Name, "\\[", "`-/", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	Name = strings.Replace(Name, "\\]", "`+/", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	arr := strings.Split(Name, ".")                //先分割路径
	arrLen := len(arr)                             //取出路径数量
	if arrLen < 1 {                                //如果数量小于1，直接返回
		return
	}
	if arrLen == 1 { //如果路径是顶级路径直接赋值
		delete(qt_json.json, arr[0])
		return
	}
	var tmpjson []interface{}       //初始化一个数组来存放原始数据
	for i := 0; i < arrLen-1; i++ { //因为最后一个的值的昵称，所以需要-1
		tmpjson = append(tmpjson, nil) //定义数组长度
	}
	str := ""
	for i := 0; i < arrLen-1; i++ {
		str = str + arr[i] + "."                         //构建路径结构
		tmp := qt_json.get_All_type(str[0 : len(str)-1]) //在原始数据中获取路径值
		switch vv := tmp.(type) {                        //判断得到的值的类型
		case map[string]interface{}: //如果是路径结构就不操作
			tmpjson[i] = vv
		case []map[string]interface{}: //如果是路径数组结构就不操作
			tmpjson[i] = vv
		case []interface{}: //如果是数组结构也不操作
			tmpjson[i] = vv
		case nil: //如果是nil 直接返回，因为没有这里路径
			return
		default: //如果 直接返回，不支持删除操作
			if qt_json.is_type {
				fmt.Println(reflect.TypeOf(vv), "DEL ")
			}
			return
		}
	}
	for i := 0; i < arrLen; i++ {
		arr[i] = txtReplace(arr[i]) //还原路径中的.
	}
	tmpname := (arr[arrLen-1]) //将值赋给最后一个路径
	stmp := tmpjson[arrLen-2]
	switch vv := stmp.(type) {
	case map[string]interface{}: //如果是路径结构就不操作
		delete(vv, tmpname)
		tmpjson[arrLen-2] = vv
	case []interface{}: //如果是数组结构也不操作
		num := toint(tmpname)
		if num == -1 {
			return
		}
		if num > len(vv)-1 {
			return
		}
		vv = append(vv[:num], vv[num+1:]...)
		tmpjson[arrLen-2] = vv
	}
	qt_json.还原结构体(arrLen, arr, tmpjson)
}

func (qt_json *qtJson) 还原结构体(arrLen int, arr []string, tmpjson []interface{}) {
	for i := 0; i < arrLen-1; i++ { //倒序还原
		x := arrLen - i - 2 //倒序还原
		if x == 0 {
			qt_json.json[(arr[x])] = tmpjson[x] //还原到结构体
		} else {
			stmp := tmpjson[x]
			switch vv := stmp.(type) {
			case map[string]interface{}: //还原到结构体
				num := toint(arr[x])
				if num == -1 {
					vq := make(map[string]interface{})
					vq[arr[x]] = tmpjson[x] //还原到结构体
					tmpjson[x-1] = vq
				} else {
					sd := tmpjson[x-1]
					switch vvs := sd.(type) {
					case map[string]interface{}: //还原到结构体
						var t []map[string]interface{}
						for o := 0; o < (num+1+len(vv))-len(vv); o++ {
							t = append(t, make(map[string]interface{}))
						}
						t[num] = vv
						sd = t //还原到结构体
						tmpjson[x-1] = sd
					case []interface{}:
						if num > 0 && num > (len(vvs)-1) {
							for i := 0; i < num-1; i++ {
								vvs = append(vvs, nil)
							}
						}
						if len(vvs)-1 < num {
							vvs = append(vvs, vv)
						} else {
							vvs[num] = vv
						}
						tmpjson[x-1] = vvs
					case []map[string]interface{}:
						if num > 0 && num > (len(vvs)-1) {
							for i := 0; i < num; i++ {
								vvs = append(vvs, nil)
							}
						}
						if len(vvs)-1 < num {
							vvs = append(vvs, vv)
						} else {
							vvs[num] = vv
						}
						tmpjson[x-1] = vvs
					default:
						if qt_json.is_type {
							fmt.Println(reflect.TypeOf(vvs), "还原到结构体1")
						}
					}
				}
			case []interface{}: //还原到结构体
				stmps := tmpjson[arrLen-3]
				switch vvs := stmps.(type) {
				case map[string]interface{}: //还原到结构体
					vvs[arr[x]] = vv
					tmpjson[arrLen-3] = vvs
				}
			}

		}
	}
}

//设置任意类型的值
func (qt_json *qtJson) set_all_type(Name string, val interface{}) {
	Name = strings.Replace(Name, "\\.", "\\+", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	Name = strings.Replace(Name, "\\[", "`-/", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	Name = strings.Replace(Name, "\\]", "`+/", -1) //暂时将 \.  换成\+ 后面真正用的时候还原
	//jsontmp := qt_json.json
	arr := strings.Split(Name, ".") //先分割路径
	arrLen := len(arr)              //取出路径数量
	if arrLen < 1 {                 //如果数量小于1，直接返回
		return
	}
	if arrLen == 1 { //如果路径是顶级路径直接赋值
		if Name == "" {
			switch vv := val.(type) {
			case map[string]interface{}: //如果是路径结构就不操作
				qt_json.json = vv
			}
		} else {
			qt_json.json[Name] = val
		}
		return
	}
	var tmpjson []interface{}       //初始化一个数组来存放原始数据
	for i := 0; i < arrLen-1; i++ { //因为最后一个的值的昵称，所以需要-1
		tmpjson = append(tmpjson, nil) //定义数组长度
	}
	str := ""
	for i := 0; i < arrLen-1; i++ {
		str = str + arr[i] + "."                         //构建路径结构
		tmp := qt_json.get_All_type(str[0 : len(str)-1]) //在原始数据中获取路径值
		switch vv := tmp.(type) {                        //判断得到的值的类型
		case map[string]interface{}: //如果是路径结构就不操作
			tmpjson[i] = vv
		case []map[string]interface{}: //如果是路径数组结构就不操作
			tmpjson[i] = vv
		case []interface{}: //如果是数组结构也不操作
			tmpjson[i] = vv
		default:
			if i+1 < len(arr) {
				inttmp := toint(arr[i+1])
				if inttmp == -1 {
					tmpjson[i] = make(map[string]interface{}) //如果不是路径结构，就重新定义为路径结构
				} else {
					var vtmps []interface{}
					tmpjson[i] = vtmps //如果不是路径结构，就重新定义为路径结构
				}
			} else {
				tmpjson[i] = make(map[string]interface{}) //如果不是路径结构，就重新定义为路径结构
			}

			/* if qt_json.is_type {
				fmt.Println(reflect.TypeOf(vv), "set_all_type ")
			} */

		}
	}
	for i := 0; i < arrLen; i++ {
		arr[i] = txtReplace(arr[i]) //还原路径中的.
	}
	tmpname := (arr[arrLen-1]) //将值赋给最后一个路径
	stmp := tmpjson[arrLen-2]
	switch vv := stmp.(type) {
	case map[string]interface{}: //如果是路径结构就不操作
		vv[tmpname] = val //赋值  因为是从0开始，最后一个是值的昵称，所以-2
		tmpjson[arrLen-2] = vv
	case []interface{}: //如果是数组结构也不操作
		num := toint(tmpname)
		nums := num - (len(vv) - 1)
		if nums > 0 && num > (len(vv)-1) {
			for i := 0; i < nums; i++ {
				vv = append(vv, 0)
			}
		}
		vv[num] = val
		tmpjson[arrLen-2] = vv
	}
	qt_json.还原结构体(arrLen, arr, tmpjson)
}

//获取任意类型的值
func (qt_json *qtJson) get_All_type(Name string) interface{} {
	arr := strings.Split(strings.Replace(strings.Replace(strings.Replace(Name, "\\[", "`-/", -1), "\\]", "`+/", -1), "\\.", "\\+", -1), ".") //暂时将 \.  换成\+ 后面真正用的时候还原

	arrLen := len(arr)
	if arrLen < 1 {
		return qt_json.json
	}
	if arr[0] == "" {
		return qt_json.json
	}
	v := qt_json.json[arr[0]] //先获取第一个路径下的对象
	if arrLen > 1 {           //如果需要下级目录
		for i := 1; i < arrLen; i++ { //因为在上一行代码获取了第一个路径 ，所以这里从1开始
			v = qt_json.get(v, arr[i]) //在上个对象中寻找下个路径
			if v == nil {              //如果获取到的是空对象 则直接返回 nil 不再继续获取
				return v
			}
		}
	}
	return v //不管是什么类型，先直接返回
}
func (qt_json *qtJson) get(val interface{}, Name string) interface{} {
	Name = strings.Replace(Name, "\\+", ".", -1) //暂时将 \+  换成.
	Name = strings.Replace(Name, "`+/", "]", -1) //暂时将 \+  换成.
	Name = strings.Replace(Name, "`-/", "[", -1) //暂时将 \+  换成.
	switch vv := val.(type) {
	case string:
		return vv
	case float64:
		return vv
	case float32:
		return vv
	case int:
		return vv
	case int8:
		return vv
	case int16:
		return vv
	case int32:
		return vv
	case int64:
		return vv
	case nil:
		return vv
	case map[string]interface{}:
		return vv[Name]
	case []map[string]interface{}:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []string:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []interface{}:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case map[int]interface{}:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []int:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []int8:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []int16:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []int32:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	case []int64:
		num := tmpyuejie(len(vv), Name)
		if num == -1 {
			return nil
		}
		return vv[num]
	default:
		if qt_json.is_type {
			fmt.Println(reflect.TypeOf(vv), "get")
		}
	}
	return nil
}
func tmpyuejie(val int, Name string) (num int) {
	if len(Name) < 3 {
		return -1
	}
	if Name[:1] != "[" || Name[len(Name)-1:len(Name)] != "]" {
		return -1
	}
	tmp, err := strconv.Atoi(Name[1 : len(Name)-1])
	if err != nil {
		return -1
	}
	if val-1 < tmp { //检查是否越界
		return -1
	}
	return tmp
}
func toint(Name string) (num int) {
	if len(Name) < 3 {
		return -1
	}
	if Name[:1] != "[" || Name[len(Name)-1:len(Name)] != "]" {
		return -1
	}
	tmp, err := strconv.Atoi(Name[1 : len(Name)-1])
	if err != nil {
		return -1
	}
	if err != nil {
		return -1
	}
	return tmp
}

//将Json格式化为String
//
//如果需要格式化 请调用 Format(true)
func (qt_json *qtJson) ToString() string {
	var content string
	if qt_json.format {
		data, err := json.MarshalIndent(qt_json.json, "", "      ")
		if err != nil {
			qt_json.err = err.Error()
			return ""
		}
		content = string(data)
	} else {
		data, err := json.Marshal(qt_json.json)
		if err != nil {
			qt_json.err = err.Error()
			return ""
		}
		content = string(data)
	}
	content = strings.Replace(content, "\\u003c", "<", -1)
	content = strings.Replace(content, "\\u003e", ">", -1)
	content = strings.Replace(content, "\\u0026", "&", -1)
	return content
}

//解析Json字符串
//
//失败返回错误信息
//
//成功返回nil
func (qt_json *qtJson) Untie(json_str string) error {
	var tmpjson interface{}
	Boby := []byte(json_str)
	err1 := json.Unmarshal(Boby, &tmpjson)
	if err1 != nil {
		qt_json.err = err1.Error()
		return err1
	}
	qt_json.json = tmpjson.(map[string]interface{})
	return nil
}
func newJson() *qtJson {
	tmp := new(qtJson)
	tmp.Untie("{}")
	return tmp
}
