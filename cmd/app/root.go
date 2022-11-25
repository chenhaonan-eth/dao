package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd表示在没有任何子命令的情况下的基本命令
var rootCmd = &cobra.Command{
	// Command的用法，Use是一个行用法消息
	Use: "grpc",
	// Short是help命令输出中显示的简短描述
	Short: "Run the gRPC Dao server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
