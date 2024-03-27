package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/ppdxzz/go-holiday/holiday"
	"github.com/robfig/cron/v3"
	"time"
)

func Job(jobGroups []*openwechat.Group) {
	fmt.Println("定时任务启动->", time.Now().Format("2006-01-02 15:04:05"))
	c := cron.New()
	// Task1
	_, err1 := c.AddFunc("0 10,15,17,22 * * *", func() {
		date := time.Now().Format("2006-01-02")
		if isWeekday, _ := holiday.IsWeekday(date); isWeekday {
			for _, group := range jobGroups {
				sendText(group, "「饮水提醒」朋友们，喝水时间到了呀，请及时喝水。")
			}
		}
	})
	if err1 != nil {
		fmt.Println(err1)
	}

	// Task2
	_, err2 := c.AddFunc("30 8 * * *", func() {
		for _, group := range jobGroups {
			sendText(group, ZaoAn("ZA"))
		}
	})
	if err2 != nil {
		fmt.Println(err2)
	}

	// Task3
	_, err3 := c.AddFunc("50 22 * * *", func() {
		for _, group := range jobGroups {
			sendText(group, ZaoAn("WA"))
		}
	})
	if err3 != nil {
		fmt.Println(err3)
	}

	// Task4
	_, err4 := c.AddFunc("30 22 * * *", func() {
		for _, group := range jobGroups {
			sendText(group, Weather("北京", true))
		}
	})
	if err4 != nil {
		fmt.Println(err4)
	}

	c.Start()
}
