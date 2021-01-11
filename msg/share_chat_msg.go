package msg

/* 群名片
{
	"msg_type": "share_chat",
	"content":{
			"share_chat_id": "oc_f5b1a7eb27ae2c7b6adc2a74faf339ff"
	}
} 
*/
type ShareChatMsg struct {
	*Sign
	MsgType string           `json:"msg_type,omitempty"`
	Content ShareChatContent `json:"content,omitempty"`
}

type ShareChatContent struct {
	ShareChatID string `json:"share_chat_id,omitempty"`
}

func NewShareChatMsg(chatID string) *ShareChatMsg{
	return &ShareChatMsg{
		MsgType: "share_chat",
		Content: ShareChatContent{
			ShareChatID: chatID,
		},
	}
}

func NewShareChatMsgWithSign(secret, chatID string) *ShareChatMsg{
	sign := NewSign(secret)
	return &ShareChatMsg{
		Sign: sign,
		MsgType: "share_chat",
		Content: ShareChatContent{
			ShareChatID: chatID,
		},
	}
}