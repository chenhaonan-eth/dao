package utils

import (
	"strings"
	"time"
	"unicode"

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

// 上月最后一天
func LastDayOfLastMonth() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	return end
}

// 解析时间
func ParseTime(datestr string) (time.Time, error) {
	for _, v := range datestr {
		if unicode.Is(unicode.Han, v) {
			str := string(v)
			if str == "日" {
				datestr = strings.Replace(datestr, str, "", -1)
				continue
			}
			datestr = strings.Replace(datestr, str, "/", -1)
		}
	}
	return dateparse.ParseAny(datestr)
}

// 中文index
func IsChinese(str string) []int {
	// var count int
	indexs := make([]int, 0)
	for i, v := range str {
		if unicode.Is(unicode.Han, v) {
			indexs = append(indexs, i)
		}
	}
	return indexs
}
