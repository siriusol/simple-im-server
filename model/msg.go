package model

import "github.com/siriusol/simple-im-server/common"

const (
	MessageTypeText = "text"
	MessageCmdMsg   = "msg"
	MessageCmdEnter = "enter"
	MessageCmdExit  = "exit"
)

type Message struct {
	Target string `json:"target"`
	Type   string `json:"type"`
	Msg    string `json:"msg"`
	From   string `json:"from"`
}

func NewTextMsg(from string, msg string) (message *Message) {
	message = &Message{
		Type: MessageTypeText,
		From: from,
		Msg:  msg,
	}
	return
}

func getTextMsgData(cmd, uuId, msgId, message string) string {
	textMsg := NewTextMsg(uuId, message)
	head := NewResponseHead(msgId, cmd, common.OK, textMsg)
	return head.String()
}

func GetMsgData(uuid, msgId, cmd, message string) string {
	return getTextMsgData(cmd, uuid, msgId, message)
}

func GetTextMsgData(uuid, msgId, message string) string {
	return getTextMsgData("msg", uuid, msgId, message)
}

func GetTextMsgDataEnter(uuid, msgId, message string) string {
	return getTextMsgData("enter", uuid, msgId, message)
}

func GetTextMsgDataExit(uuid, msgId, message string) string {
	return getTextMsgData("exit", uuid, msgId, message)
}
