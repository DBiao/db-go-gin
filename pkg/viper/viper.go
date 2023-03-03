package viper

import (
	"db-go-gin/internal/global"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitViper(path ...string) (*viper.Viper, error) {
	var config string
	if len(path) == 0 || path[0] == "" {
		flag.StringVar(&config, "path", "", "choose config file.")
		flag.Parse()

		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv("BACK_CONFIG"); configEnv == "" {
				config = "./conf/config.yaml"
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用BACK_CONFIG环境变量,config的路径为%v\n", config)
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
		return v, err
	}
	v.WatchConfig()

	// 文件被修改时打印日志
	v.OnConfigChange(func(e fsnotify.Event) {
		global.LOG.Info("config file changed:", zap.String("name", e.Name))
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			global.LOG.Error("config file changed error", zap.Error(err))
		}
	})

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		return v, err
	}

	return v, nil
}
