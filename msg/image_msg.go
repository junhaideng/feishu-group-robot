package msg

/* 图片
{
    "msg_type":"image",
    "content":{
        "image_key": "img_ecffc3b9-8f14-400f-a014-05eca1a4310g"
    }
}
*/

type ImageMsg struct {
	*Sign
	MsgType string       `json:"msg_type,omitempty"`
	Content ImageContent `json:"content,omitempty"`
}

type ImageContent struct {
	ImageKey string `json:"image_key,omitempty"`
}

func NewImageMsg(imageKey string) *ImageMsg {
	return &ImageMsg{
		MsgType: "image",
		Content: ImageContent{
			ImageKey: imageKey,
		},
	}
}

func NewImageMsgWithSign(secret, imageKey string) *ImageMsg {
	sign := NewSign(secret)
	return &ImageMsg{
		Sign:    sign,
		MsgType: "image",
		Content: ImageContent{
			ImageKey: imageKey,
		},
	}
}
