package model

import (
	"encoding/json"
	"github.com/siriusol/simple-im-server/common"
)

type Head struct {
	Seq      string    `json:"seq"`
	Cmd      string    `json:"cmd"`
	Response *Response `json:"response"` // 消息体
}

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PushMsg struct {
	Seq  string `json:"seq"`
	UUID int64  `json:"uuid"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

// NewResponseHead 设置返回消息
func NewResponseHead(seq string, cmd string, code common.HTTPError, data interface{}) *Head {
	response := NewResponse(code, data)
	return &Head{
		Seq:      seq,
		Cmd:      cmd,
		Response: response,
	}
}

func (h *Head) String() string {
	headBytes, _ := json.Marshal(h)
	return string(headBytes)
}

func NewResponse(code common.HTTPError, data interface{}) *Response {
	return &Response{
		Code: code.Code(),
		Msg:  code.Msg(),
		Data: data,
	}
}
