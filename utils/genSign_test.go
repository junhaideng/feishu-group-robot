package utils

import (
	"testing"
)

func TestGenSign(t *testing.T){
	timeStamp := "1609773929"
	println(GenSign("Z3XBKDsqRepLGr1KP3aarb", timeStamp))
}