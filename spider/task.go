package spider

import (
	"github.com/chenhaonan-eth/dao/config"
	cron "github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func StartTask() {
	crontab := cron.New()
	task := func() {
		defer func() {
			if err := recover(); err != nil {
				config.G_LOG.Error("task ", zap.Any("err", err))
			}
		}()
		CollyCNPPI()
		CollyCNCPI()
		CollyCNSocialFinancingStock()
		CollyChinaMoneySupply()
		CollyMacroChinaConsumerGoodsRetail()
		CollyCNSocialFinancingFlow()
		CollyValueAddedOfIndustrialProduction()
		CollySocialElectricityConsumption()
	}

	FuturesForeignTask := func() {
		defer func() {
			if err := recover(); err != nil {
				config.G_LOG.Error("task ", zap.Any("err", err))
			}
		}()
		CollyCADFuturesForeignHist()
		CollySH300PE()
		CollyBondZhUsRate()
	}

	// 月初第一天10点
	crontab.AddFunc("0 10 1 * *", func() {
		defer func() {
			if err := recover(); err != nil {
				config.G_LOG.Error("task ", zap.Any("err", err))
			}
		}()
		CollyCNPMI()
		CollyCXPMI()
	})
	//每月15-30日上午10:00触发
	crontab.AddFunc("0 10 15-30 * *", task)
	// 每日9点
	crontab.AddFunc("0 9 * * *", FuturesForeignTask)
	//1月、4月、7月、10月17日-19 9点半
	crontab.AddFunc("0 9 17-19 1,4,7,10 *", func() {
		defer func() {
			if err := recover(); err != nil {
				config.G_LOG.Error("task ", zap.Any("err", err))
			}
		}()
		CollyCNGDP()
	})

	// 启动定时器
	crontab.Start()
}
