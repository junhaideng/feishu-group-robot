package msg

/*
{
    "msg_type": "interactive",
    "card": {
        "config": {
                "wide_screen_mode": true,
                "enable_forward": true
        },
        "elements": [{
                "tag": "div",
                "text": {
                        "content": "**西湖**，位于浙江省杭州市西湖区龙井路1号，杭州市区西部，景区总面积49平方千米，汇水面积为21.22平方千米，湖面面积为6.38平方千米。",
                        "tag": "lark_md"
                }
        }, {
                "actions": [{
                        "tag": "button",
                        "text": {
                                "content": "更多景点介绍 :玫瑰:",
                                "tag": "lark_md"
                        },
                        "url": "https://www.example.com",
                        "type": "default",
                        "value": {}
                }],
                "tag": "action"
        }],
        "header": {
                "title": {
                        "content": "今日旅游推荐",
                        "tag": "plain_text"
                }
        }
    }
}
*/

type CardMsg struct {
	*Sign
	MsgType string `json:"msg_type,omitempty"`
	Card    Card   `json:"card,omitempty"`
}


type Card struct {
	Config   Config    `json:"config,omitempty"`
	Elements []Element `json:"elements,omitempty"`
	Header   Header    `json:"header,omitempty"`
}

type Element struct {
	Tag  string `json:"tag,omitempty"`
	Text Body   `json:"text,omitempty"`
}


type Config struct {
	WideScreenMode bool `json:"wide_screen_mode,omitempty"`
	EnableForward  bool `json:"enable_forward,omitempty"`
}

type Header struct {
	Title Body `json:"title,omitempty"`
}

type Body struct {
	Content string `json:"content,omitempty"`
	Tag     string `json:"tag,omitempty"`
}



func NewCardMsg(title string) *CardMsg{
	return &CardMsg{
		MsgType: "interactive",
		Card: Card{
			Config: Config{
				WideScreenMode: true,
				EnableForward: true,
			},
			Header: Header{
				Title: Body{
					Tag: "plain_text",
					Content: title,
				},
			},
		},
	}
}


func NewCardMsgWithSign(secret, title string) *CardMsg{
	sign := NewSign(secret)
	return &CardMsg{
		Sign: sign,
		MsgType: "interactive",
		Card: Card{
			Config: Config{
				WideScreenMode: true,
				EnableForward: true,
			},
			Header: Header{
				Title: Body{
					Tag: "plain_text",
					Content: title,
				},
			},
		},
	}
}

// 添加一个内容，内容的格式为markdown形式
func(c *CardMsg)AddElement(content string){
	element := Element{
		Tag: "div",
		Text: Body{
			Content: content, 
			Tag: "lark_md",
		},
	}
	c.Card.Elements = append(c.Card.Elements, element)
}