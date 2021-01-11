package msg

import (
	"encoding/json"
	"testing"
)

func TestCardMsg(t *testing.T){
	msg := NewCardMsg("just a test")
	msg.AddElement("**bold content**")
	data, err := json.MarshalIndent(msg, " ", "  ")
	if err != nil{
		t.Fatal(err)
	}
	t.Log(string(data))
}