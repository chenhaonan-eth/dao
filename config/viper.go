package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	G_Config      Config
	G_viper       *viper.Viper
	ConfigEnvPath string
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
		if ConfigEnvPath == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv("DAO_CONFIG"); configEnv == "" {
				config = "config.yaml"
				fmt.Printf("You are using the default value of config, whose path is:%v\n", "config.yaml")
			} else {
				config = configEnv
				fmt.Printf("You are using the DAO CONFIG environment variable, and the path to config is:%v\n", config)
			}
		} else {
			config = ConfigEnvPath
			fmt.Printf("The value you are passing using the -c argument on the command line, the path to config is:%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("The value you are passing using func Viper(), the path of config is:%v\n", config)
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
