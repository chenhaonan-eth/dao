package router

import (
	"github.com/chenhaonan-eth/guide-sisyphean/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	router := gin.Default()

	WebInit(router)

	if err := router.Run(core.G_Config.System.Port); err != nil {
		core.G_LOG.Error("router init fail", zap.Error(err))
	}
}
