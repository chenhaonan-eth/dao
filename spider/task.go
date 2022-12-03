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
		CollyMacroChinaConsumerGoodsRetail()
		CollyCNSocialFinancingFlow()
	}

	FuturesForeignTask := func() {
		CollyCADFuturesForeignHist()
		CollySH300PE()
		CollyBondZhUsRate()
	}

	// 月初第一天10点半
	crontab.AddFunc("0 15 10 1 * ?", func() { CollyCNPMI() })
	//每月9-15日上午10:15触发
	crontab.AddFunc("0 15 10 9-15 * ?", task)
	// 每日9点
	crontab.AddFunc("0 15 9 * * ?", FuturesForeignTask)
	//1月、4月、7月、10月17日-19 9点半
	crontab.AddFunc("0 15 9 17-19 1,4,7,10 ?", func() { CollyCNGDP() })
	// 启动定时器
	crontab.Start()
}
