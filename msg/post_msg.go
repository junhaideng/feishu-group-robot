package msg

/* 富文本消息	
{
    "msg_type": "post",
    "content": {
        "post": {
            "zh_cn": {
                "title": "项目更新通知",
                "content": [
                    [
                        {
                            "tag": "text",
                            "text": "项目有更新: "
                        },
                        {
                            "tag": "a",
                            "text": "请查看",
                            "href": "http://www.example.com/"
                        }
                    ]
                ]
            }
        }
    }
} 
*/

type PostMsg struct {
	*Sign
	MsgType string      `json:"msg_type,omitempty"`
	Content PostContent `json:"content,omitempty"`
}

type PostContent struct {
	Post Post `json:"post,omitempty"`
}

type Post struct {
	ZhCn ZhCn `json:"zh_cn,omitempty"`
}

type ZhCn struct {
	Title   string `json:"title,omitempty"`
	Content [][]Tag  `json:"content,omitempty"`
}

type Tag struct {
	// text or a 
	Tag  string `json:"tag,omitempty"`
	// 文本消息
	Text string `json:"text,omitempty"`
	// 如果Tag为a，需要指定超链接
	Href string `json:"href,omitempty"`
}

func NewPostMsg(title string)*PostMsg{
	return &PostMsg{
		MsgType: "post",
		Content: PostContent{
			Post: Post{
				ZhCn:ZhCn{
					Title: title, 
					Content: make([][]Tag, 1),
				},
			},
		},
	}
}

func NewPostMsgWithSign(secret, title string)*PostMsg{
	sign := NewSign(secret)
	return &PostMsg{
		Sign: sign,
		MsgType: "post",
		Content: PostContent{
			Post: Post{
				ZhCn:ZhCn{
					Title: title, 
					Content: make([][]Tag, 1),
				},
			},
		},
	}
}

// 添加普通文本
func(p *PostMsg) AddText(text string){
	content := p.Content.Post.ZhCn.Content
	c := Tag{
		Tag: "text",
		Text: text,
	}
	content[0] = append(content[0], c)
}

// 添加超链接
func(p *PostMsg) AddLink(text, href string){
	content := p.Content.Post.ZhCn.Content
	c := Tag{
		Tag: "a",
		Text: text,
		Href: href,
	}
	content[0] = append(content[0], c)
}

