package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
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
	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		handle(msg)
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 登录
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}
	// 获取登录的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化配置
	initConfig()
	// 获取所有的群组
	groups, _ := self.Groups()
	var jobGroups []*openwechat.Group
	for _, group := range groups {
		fmt.Printf("%s\n", group.NickName)
		if group.NickName == botGroupName {
			jobGroups = append(jobGroups, group)
		}
	}
	// 	启动定时任务
	Job(jobGroups)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()

}
