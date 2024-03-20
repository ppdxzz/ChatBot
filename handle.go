package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"math/rand/v2"
	"strings"
	"time"
)

func handle(msg *openwechat.Message) {
	if !msg.IsText() {
		return
	}
	if !msg.IsSendBySelf() && msg.IsSendByGroup() {
		handleGroup(msg)
	} else if msg.IsSendByFriend() {
		handleFriend(msg)
	}
}

// 处理群组消息
func handleGroup(msg *openwechat.Message) {
	sender, _ := msg.Sender()
	group := openwechat.Group{User: sender}
	if group.NickName == botGroupName {
		if strings.Contains(msg.Content, "你好") || strings.Contains(msg.Content, "是谁") {
			replyText(msg, "你好，我是小弟，三岁半的小孩子，请不要欺负我哦")
		} else if strings.HasSuffix(msg.Content, "天气") && len(msg.Content) > 2 {
			replyText(msg, Weather(strings.Replace(msg.Content, "天气", "", -1), false))
		} else if strings.HasPrefix(msg.Content, "小弟") && len(msg.Content) > 2 {
			go replyText(msg, Chat(strings.Replace(msg.Content, "小弟", "", -1)))
		} else if msg.IsJoinGroup() {
			replyText(msg, "哟，新人进群了，大家欢迎呀")
		}
	}
}

// 处理私聊消息
func handleFriend(msg *openwechat.Message) {
	if msg.Content == "ping" {
		replyText(msg, "pong")
	}
}

func replyText(msg *openwechat.Message, content string) {
	time.Sleep(time.Second * time.Duration(rand.IntN(6)+3))
	_, err := msg.ReplyText(content)
	if err != nil {
		fmt.Println("回复消息异常：", err)
	}
}

func replyTextAndRevoke(msg *openwechat.Message, content string) {
	sentMessage, err := msg.ReplyText(content)
	if err != nil {
		fmt.Println(err)
	}
	go revokeMessage(sentMessage)
}

func revokeMessage(sentMessage *openwechat.SentMessage) {
	time.Sleep(time.Second * time.Duration(rand.IntN(11)+10))
	if sentMessage.CanRevoke() {
		_ = sentMessage.Revoke()
	}
}

func sendText(group *openwechat.Group, content string) {
	time.Sleep(time.Second * time.Duration(rand.IntN(6)+3))
	_, err := group.SendText(content)
	if err != nil {
		fmt.Println("发送消息异常：", err)
	}
}
