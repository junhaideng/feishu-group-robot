## 飞书群机器人

> 参考 [飞书文档](https://www.feishu.cn/hc/zh-cn/articles/360024984973-在群聊中使用机器人) 进行简单实现，能够简单的业务功能，适合用来进行消息通知

#### 下载
```
go get github.com/junhaideng/feishu-group-robot
```

#### 示例代码
```go
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

// 使用签名验证的图片消息
func sendImageMsgWithSign() {
	m := msg.NewImageMsgWithSign(secret, "img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	fmt.Println(utils.SendMsg(webhook, m))
}

func sendShareChatMsg() {
	m := msg.NewShareChatMsg("oc_f5b1a7eb27ae2c7b6adc2a74faf339ff")
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

```

#### 注意
> 使用签名验证的时候，发送消息的时候有时候可以发送成功，但是有时候发送可能会发送失败，没有签名验证的时候不存在这种问题