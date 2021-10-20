package qt

import (
	"time"
)

type Qttime struct {
}

// GeTimeString
//
// 获取字符串类型的时间文本 “2006-01-02 15:04:05”
func (c Qttime) GeTimeString() string {
	t := time.Now()
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
