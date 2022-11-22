package spider

import (
	cron "github.com/robfig/cron/v3"
)

func StartTask() {
	crontab := cron.New()
	task := func() {
		CollyCNPPI()
		CollyCNSocialFinancingStock()
	}

	//每月9-15日上午10:15触发
	crontab.AddFunc("0 15 10 9-15 *", task)

	// 启动定时器
	crontab.Start()
}
