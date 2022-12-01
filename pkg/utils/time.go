package utils

import (
	"time"
)

// 获取上个月第一天
func FirstDayOfLastMonth() string {
	const DATE_FORMAT = "2006-01-02 15:04:05"
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	return start
}

func LastDayOfLastMonth() string {
	const DATE_FORMAT = "2006-01-02 15:04:05"
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	return end
}
