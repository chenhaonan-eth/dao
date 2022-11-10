// 宏观周期指标
package api

import (
	"github.com/chenhaonan-eth/guide-sisyphean/core"
	"github.com/chenhaonan-eth/guide-sisyphean/economic"
	"github.com/chenhaonan-eth/guide-sisyphean/utils.go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetdEquityBbondYieldSpreads(c *gin.Context) {
	if v, err := economic.EquityBbondYieldSpreads(); err != nil {
		core.G_LOG.Error("find Equity Bond Yield Spreads err", zap.Any("err", err))
		utils.FailWithMessage("find Equity Bond Yield Spreads err", c)
	} else {
		utils.OkWithDetailed(v, "find Equity Bond Yield Spreads succeed", c)
	}
}
