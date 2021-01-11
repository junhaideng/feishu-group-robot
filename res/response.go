package res

import (
	"encoding/json"
)

// 返回的错误信息
type ErrResponse struct{
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

// 成功发送之后的消息
type Response struct {
	Extra         json.RawMessage
	StatusCode    int
	StatusMessage string
}
