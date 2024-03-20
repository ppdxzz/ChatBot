package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

const botGroupName = "皮皮乐开发者社群"

func initConfig() {
	viper.SetConfigName("config-dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		} else {
			panic(err)
		}
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func main() {
	// 初始化配置
	initConfig()
}
