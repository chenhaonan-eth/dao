package main

import (
	_ "github.com/chenhaonan-eth/dao/core"
	"github.com/chenhaonan-eth/dao/router"
)

// go:generate go env -w GO111MODULE=on
// go:generate go env -w GOPROXY=https://goproxy.cn,direct
// go:generate go mod tidy
// go:generate go mod download

//TODO: 启动时候检查配置文件是否正确
func main() {
	router.Run()
}