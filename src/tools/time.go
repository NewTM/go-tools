package tools

import "time"

const TimeLayout = "2016-01-02 15:04:05"

//时间字符串转化成 时间对象
func StrToTime(str string) (time.Time, error) {
	res, err := time.ParseInLocation(TimeLayout, str, time.Local)
	if err != nil {
		return time.Now(), err
	}
	return res, nil
}
