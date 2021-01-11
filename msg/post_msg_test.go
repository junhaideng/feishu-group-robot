package msg

import (
	"testing"
	"encoding/json"
)

func TestNewPostMsg(t *testing.T){
	msg := NewPostMsg("post msg")
	msg.AddText("this is a post msg from robot")
	data, err := json.MarshalIndent(msg, " ", "  ")
	if err != nil{
		t.Fatal(err)
	}
	t.Log(string(data))
}