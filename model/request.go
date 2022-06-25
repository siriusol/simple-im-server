package model

type Request struct {
	Seq  string      `json:"seq"` // 消息的唯一Id
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data,omitempty"` // 数据json
}

type Login struct {
	ServiceToken string `json:"service_token"` //验证用户是否登录
	AppId        int32  `json:"app_id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
}

type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
