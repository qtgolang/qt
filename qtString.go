package qt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type qtstring struct{}

// Scanf 从控制台获取文本输入
func (c qtstring) Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}

func (c *qtstring) Init64Tostring(val int64) string {
	return strconv.FormatInt(val, 10)
}

//Int_To_string  数值类型转字符串
func (c *qtstring) Int_To_string(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 32)
	case float32:
		return strconv.FormatFloat(float64(v), 'E', -1, 32)
	}
	return ""
}

//Float64Tostring val=操作的数值 baoli=保留小数后几位
func (c *qtstring) Float64Tostring(val float64, baoliu int) string {
	value := fmt.Sprintf("%."+strconv.Itoa(baoliu)+"f", val)
	return value
}

//StringToFloat64   String_To_Numer  字符串转数值类型
func (c *qtstring) StringToFloat64(val string) float64 {
	v, _ := strconv.ParseFloat(val, 64)
	return v
}
func (c *qtstring) Init64ToInt(val int64) int {
	int_num := c.Init64Tostring(val)
	i, _ := strconv.Atoi(int_num)
	return i
}

//SubString 取字符串中间
func (c qtstring) SubString(str, left, Right string) string {
	s := strings.Index(str, left)
	if s < 0 {
		return ""
	}
	s += len(left)
	e := strings.Index(str[s:], Right)
	if e+s <= s {
		return ""
	}
	return str[s : s+e]

}

//SubString 批量取字符串中间
func (c qtstring) SubStringArr(str, left, Right string) []string {
	var 存放取出文本的数组 []string
	StartPos := 0
	EndPos := 0
	中间长度 := 0
	Search := ""
	i := 0
	for {
		i = strings.Index(str[StartPos:], left)
		if i != -1 {
			StartPos = StartPos + i
		} else {
			StartPos = i
		}
		if StartPos != -1 {
			StartPos = StartPos + len(left)
			EndPos = StartPos + strings.Index(str[StartPos:], Right)
			if EndPos != -1 {
				中间长度 = EndPos - StartPos
				Search = str[StartPos : StartPos+中间长度]
				存放取出文本的数组 = append(存放取出文本的数组, Search)

			}

		}

		if StartPos == -1 {
			break
		}
	}

	return 存放取出文本的数组

}


//New 初始化一个qtString对象
func newstring() *qtstring {
	t := new(qtstring)
	return t
}
