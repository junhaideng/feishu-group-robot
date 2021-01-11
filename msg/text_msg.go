package msg

/* 文本消息
{
    "msg_type": "text",
    "content": {
        "text": "新更新提醒"
    }
}
*/
type TextMsg struct {
	*Sign
	MsgType string      `json:"msg_type,omitempty"`
	Content TextContent `json:"content,omitempty"`
}

type TextContent struct {
	Text string `json:"text,omitempty"`
}

func NewTextMsg(text string) *TextMsg {
	return &TextMsg{
		MsgType: "text",
		Content: TextContent{
			Text: text,
		},
	}
}

func NewTextMsgWithSign(secret, text string) *TextMsg {
	sign := NewSign(secret)
	return &TextMsg{
		Sign:    sign,
		MsgType: "text",
		Content: TextContent{
			Text: text,
		},
	}
}
