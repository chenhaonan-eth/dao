package spider

import cron "github.com/robfig/cron/v3"

func StartTask() {
	crontab := cron.New()
	task := func() {
		CollyCNPPI()
		CollyCNCPI()
		CollyCNSocialFinancingStock()
		CollyChinaMoneySupply()
		CollyMacroChinaConsumerGoodsRetail()
		CollyCNSocialFinancingFlow()
		CollyValueAddedOfIndustrialProduction()
	}

	FuturesForeignTask := func() {
		CollyCADFuturesForeignHist()
		CollySH300PE()
		CollyBondZhUsRate()
	}

	// 月初第一天10点
	crontab.AddFunc("0 10 1 * *", func() {
		CollyCNPMI()
		CollyCXPMI()
	})
	//每月9-15日上午10:00触发
	crontab.AddFunc("0 10 9-15 * *", task)
	// 每日9点
	crontab.AddFunc("0 9 * * *", FuturesForeignTask)
	//1月、4月、7月、10月17日-19 9点半
	crontab.AddFunc("0 9 17-19 1,4,7,10 *", func() { CollyCNGDP() })

	// 启动定时器
	crontab.Start()
}
