package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// 生成消息摘要
func GenSign(secret, timestamp string) string {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
