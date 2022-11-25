package main

import (
	cmd "github.com/chenhaonan-eth/dao/cmd/app"
	// _ "github.com/chenhaonan-eth/dao/core"
	// _ "github.com/chenhaonan-eth/dao/dal/initialize"
	// "github.com/chenhaonan-eth/dao/spider"
	// "github.com/chenhaonan-eth/dao/router"
)

// go:generate go env -w GO111MODULE=on
// go:generate go env -w GOPROXY=https://goproxy.cn,direct
// go:generate go mod tidy
// go:generate go mod download

// TODO: 启动时候检查配置文件是否正确
func main() {
	// step1. 启动时检查数据库，初始化所有Table、数据

	// step2 开启定时任务，定时爬取数据并存入本地数据库
	// spider.StartTask()

	// start HTTP Grpc
	cmd.Execute()

}
