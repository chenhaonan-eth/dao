package main

import (
	"github.com/chenhaonan-eth/dao/core"
	_ "github.com/chenhaonan-eth/dao/core"
	"github.com/chenhaonan-eth/dao/dal"
	"github.com/chenhaonan-eth/dao/dal/query"
	"github.com/chenhaonan-eth/dao/router"
	"github.com/chenhaonan-eth/dao/spider"
)

// go:generate go env -w GO111MODULE=on
// go:generate go env -w GOPROXY=https://goproxy.cn,direct
// go:generate go mod tidy
// go:generate go mod download

//TODO: 启动时候检查配置文件是否正确
func main() {

	// 创建本地数据库
	dal.DB = dal.ConnectDB(core.G_Config.System.Dsn).Debug()
	query.SetDefault(dal.DB)

	//TODO: 启动时检查数据库，初始化所有Table、数据
	// db.initDB()

	//TODO: 开启定时任务，定时爬取数据并存入本地数据库
	spider.StartTask()

	//TODO: HTTP
	router.Run()
	//TODO: Grpc
	// Grpc.run()
}
