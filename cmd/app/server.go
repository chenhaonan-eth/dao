package cmd

import (
	"log"

	"github.com/chenhaonan-eth/dao/config"
	"github.com/chenhaonan-eth/dao/dal/initialize"
	"github.com/chenhaonan-eth/dao/server"
	"github.com/chenhaonan-eth/dao/spider"
	"github.com/spf13/cobra"
)

// 创建附加命令
// 本地标签：在本地分配一个标志，该标志仅适用于该特定命令。
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC and Http Dao server",
	// 运行:典型的实际工作功能。大多数命令只会实现这一点；
	// 另外还有PreRun、PreRunE、PostRun、PostRunE等等不同时期的运行命令，但比较少用，具体使用时再查看亦可
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error : %v", err)
			}
		}()

		// 配置文件 LOG初始化
		config.Init()
		// step1. 启动时检查数据库，初始化所有Table、数据
		initialize.Init()
		// step2 开启定时任务，定时爬取数据并存入本地数据库
		spider.StartTask()
		// start HTTP Grpc
		server.Run()
	},
}

// 在 init() 函数中定义flags和处理配置。
func init() {
	// 我们定义了一个flag，值存储在&server.ServerPort中，长命令为--port，短命令为-p，，默认值为50052。
	// 命令的描述为server port。这一种调用方式成为Local Flags 本地标签

	serverCmd.Flags().StringVarP(&server.Network, "network", "n", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	serverCmd.Flags().StringVarP(&server.HttpPort, "http-port", "", "50053", "http port")
	serverCmd.Flags().StringVarP(&server.ServerPort, "gRPCport", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto/server", "path to the directory which contains swagger definitions")
	serverCmd.Flags().StringVarP(&config.ConfigEnvPath, "config-env-path", "c", "", "choose config file.")
	// AddCommand向这父命令（rootCmd）添加一个或多个命令
	rootCmd.AddCommand(serverCmd)
}
