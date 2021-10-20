package qt

import (
	"strconv"
	"strings"
	"time"
)

type Qttime struct {
}

//Sleep
//
// 延迟毫秒数
func (c Qttime) Sleep(Millisecond int) {
	time.Sleep(time.Millisecond * time.Duration(Millisecond))
}

// TimeDiff
//
//TimeDiff 取两个时间的时间差
func (c Qttime) TimeDiff(time1, time2 time.Time) time.Duration {
	v := time1.Unix() - time2.Unix()
	return time.Duration(v) * time.Second
}

// StringTimeDiff
//
//StringTimeDiff 取两个字符串时间文本的时间差  “2006-01-02 15:04:05”
func (c Qttime) StringTimeDiff(time1, time2 string) time.Duration {
	return c.TimeDiff(c.StringToTime(time1), c.StringToTime(time2))
}

// TimeAdd
//
// 时间加减操作
//
// unit 操作的时间单位 1=年 2=月 3=日 4=时 5=分 6=秒
//
// val 要操作的值 例如 10 或者 -10
func (c Qttime) TimeAdd(time1 time.Time, unit, val int) time.Time {
	s := ""
	switch unit {
	case 1:
		arr := strings.Split(c.TimeToString(time1), "-")
		if len(arr) == 3 {
			i, _ := strconv.Atoi(arr[0])
			s = strconv.Itoa(i+val) + "-" + arr[1] + "-" + arr[2]
			t1 := c.StringToTime(s)
			return t1
		}
	case 2:
		arr := strings.Split(c.TimeToString(time1), "-")
		if len(arr) == 3 {
			i, _ := strconv.Atoi(arr[0])
			i2, _ := strconv.Atoi(arr[1])
			ye := (i2 + val) % 12
			if ye == 0 {
				ye = 12
			}
			n := (i2 + val - ye) / 12
			if ye < 10 {
				if ye < 0 {
					n--
					ye = ye + 12
					if ye < 10 {
						s = strconv.Itoa(i+n) + "-0" + strconv.Itoa(ye) + "-" + arr[2]
					} else {
						s = strconv.Itoa(i+n) + "-" + strconv.Itoa(ye) + "-" + arr[2]
					}
				} else {
					s = strconv.Itoa(i+n) + "-0" + strconv.Itoa(ye) + "-" + arr[2]
				}

			} else {
				s = strconv.Itoa(i+n) + "-" + strconv.Itoa(ye) + "-" + arr[2]
			}
			t1 := c.StringToTime(s)
			return t1
		}
	case 3:
		s = strconv.Itoa(val*24) + "h"
		dd, _ := time.ParseDuration(s)
		return time1.Add(dd)
	case 4:
		s = strconv.Itoa(val) + "h"
		dd, _ := time.ParseDuration(s)
		return time1.Add(dd)
	case 5:
		s = strconv.Itoa(val) + "m"
		dd, _ := time.ParseDuration(s)
		return time1.Add(dd)
	case 6:
		s = strconv.Itoa(val) + "s"
		dd, _ := time.ParseDuration(s)
		return time1.Add(dd)
	}
	return time1
}

// GetTimeString
//
// 获取字符串类型的时间文本 “2006-01-02 15:04:05”
func (c Qttime) GetTimeString() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

// TimeToString
//
// 指定字符串类型的时间文本 “2006-01-02 15:04:05”
func (c Qttime) TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeCompare
//
// 两个时间进行比较 如果 time1>time2 返回True 否则返回 False
func (c Qttime) TimeCompare(time1, time2 time.Time) bool {
	return time2.Before(time1)
}

// TimeStringcompare
//
// 两个String时间字符串 进行比较 如果 time1>time2 返回True 否则返回 False
func (c Qttime) TimeStringcompare(time1, time2 string) bool {
	return c.TimeCompare(c.StringToTime(time1), c.StringToTime(time2))
}

// StringToTime
//
// 字符串类型的时间文本 “2006-01-02 15:04:05” 到时间格式
func (c Qttime) StringToTime(strtime string) time.Time {
	formatTime, _ := time.Parse("2006-01-02 15:04:05", strtime)
	return formatTime
}

// GetTimeUnix
//
//获取当前时间戳 参数若为 true 获取10 位的时间戳，否则为13位时间戳
func (c Qttime) GetTimeUnix(val bool) int64 {
	if val {
		return time.Now().UnixNano() / 1000 / 1000 / 1000
	}
	return time.Now().UnixNano() / 1000 / 1000
}

// TimeStrToTimeUnix
//
//指定时间到时间戳 10位
func (c Qttime) TimeStrToTimeUnix(times string) int64 {

	formatTime, err := time.Parse("2006-01-02 15:04:05", times)
	if err == nil {
		return formatTime.Unix()
	}
	return -1
}
