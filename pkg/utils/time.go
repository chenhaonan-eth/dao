package utils

import (
	"time"

	"github.com/noaway/dateparse"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

// 获取上个月第一天
// return "2006-01-02 00:00:00"
func FirstDayOfLastMonth() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	return start
}

func LastDayOfLastMonth() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	return end
}

// 解析时间
func ParseTime(datestr string) (time.Time, error) {
	return dateparse.ParseAny(datestr)
}
