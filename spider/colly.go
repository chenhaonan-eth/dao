package spider

import (
	"log"
	"time"
)

// 爬取CN社融ppi，从每月9号开始
func CollyCNPPI() {
	// TODO
	// 计算本月时间
	t := time.Now().Format("2006-01")
	log.Println(t)
	// 查询数据库是否存在最近一个月数据
	// 无，爬最新数据
}

// 社会融资存量 每月10号
func CollyCNSocialFinancingStock() {
	// TODO
	// 计算本月时间
	// 查询数据库是否存在最近三个月数据
	// 无，爬最新数据

}

// 伦铜 每日
func CollyCADFuturesForeignHist() {
	// TODO
	// 计算本日时间
	// 查询数据库是否存在最近三个月数据
	// 无，爬最新数据

}
