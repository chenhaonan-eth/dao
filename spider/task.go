package spider

import (
	cron "github.com/robfig/cron/v3"
)

func StartTask() {
	crontab := cron.New()
	task := func() {
		CollyCNPPI()
		CollyCNCPI()
		CollyCNSocialFinancingStock()
		CollyChinaMoneySupply()
	}

	FuturesForeignTask := func() {
		CollyCADFuturesForeignHist()
		CollySH300PE()
	}

	// 每日任务
	// dailyTask := func (){CollySH300PE()}

	//每月9-15日上午10:15触发
	crontab.AddFunc("0 15 10 9-15 *", task)
	// 每日9点
	crontab.AddFunc("0 15 9 * *", FuturesForeignTask)
	// 启动定时器
	crontab.Start()
}
