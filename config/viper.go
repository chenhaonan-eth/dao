package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	G_Config Config
	G_viper  *viper.Viper
)

func Init() {
	G_viper = initViper()
	G_LOG = initZap()
}

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func initViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv("KASEYA_CONFIG"); configEnv == "" {
				config = "config.yaml"
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", "config.yaml")
			} else {
				config = configEnv
				fmt.Printf("您正在使用KASEYA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&G_Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&G_Config); err != nil {
		fmt.Println(err)
	}
	return v
}
