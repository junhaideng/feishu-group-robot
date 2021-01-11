package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/junhaideng/feishu-group-robot/res"
	"io/ioutil"
	"net/http"
)

func SendMsg(webhook string,v interface{}) (response interface{},isErrResponse bool,err error){
	if webhook == ""{
		return nil, false, errors.New("no specified webhook")
	}
	
	data, err := json.Marshal(v)	
	if err != nil{
		return nil, false, errors.New("marshal data error")
	}

	resp, err := http.Post(webhook, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	// 保存在[]byte中复用
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, false, err
	}

	// 返回的是成功的提示
	var res1 res.Response
	err = json.Unmarshal(d, &res1)
	if err != nil{
		return nil, false, err
	}
	if res1.StatusMessage != ""{
		return res1, false, nil
	}

	// 返回的是错误信息
	var res2 res.ErrResponse
	err = json.Unmarshal(d, &res2)
	if err != nil{
		return nil, true, err
	}
	return res2, true, nil
}
