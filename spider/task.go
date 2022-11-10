package spider

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
)

func StartTask() {
	crontab := cron.New()
	task := func() {
		fmt.Println("hello world")
	}
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	crontab.AddFunc("* * * * *", task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
}
