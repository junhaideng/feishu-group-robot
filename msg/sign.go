package msg

import (
	"github.com/junhaideng/feishu-group-robot/utils"
	"strconv"
	"time"
)

// 签名
type Sign struct {
	Timestamp string `json:"timestamp,omitempty"`
	Sign      string `json:"sign,omitempty"`
}


func NewSign(secret string) *Sign {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sign := utils.GenSign(secret, timestamp)
	return &Sign{
		Timestamp: timestamp,
		Sign:      sign,
	}
}
