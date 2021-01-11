package main

import (
	"github.com/junhaideng/feishu-group-robot/msg"
	"github.com/junhaideng/feishu-group-robot/utils"
	"fmt"
	"os"
)

// webhook url
var webhook string

// 数字签名
var secret string

func init() {
	webhook = os.Getenv("FEISHU_WEBHOOK")
	secret = os.Getenv("FEISHU_SECRET")
}

func main() {
  sendCardMsg()
}

// 发送图片消息
func sendImageMsg() {
	m := msg.NewImageMsg("img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendImageMsgWithSign() {
	m := msg.NewImageMsgWithSign(secret, "img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	fmt.Println(utils.SendMsg(webhook, m))
}

// 发送
func sendShareChatMsg() {
	m := msg.NewShareChatMsg("oc_111")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendPostMsg() {
  m := msg.NewPostMsg("post msg")
  m.AddText("hi, welcome ")
  m.AddLink("here", "https://github.com/junhaideng")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendTextMsg() {
	m := msg.NewTextMsg("hello world")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendCardMsg() {
  m := msg.NewCardMsg("just a test")
  m.AddElement("a pure **test**, use any markdown syntax you like")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendTextMsgWithSign() {
	m := msg.NewTextMsgWithSign(secret, "hello world")
	fmt.Println(utils.SendMsg(webhook, m))
}
