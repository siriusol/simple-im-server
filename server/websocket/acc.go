package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/siriusol/simple-im-server/common"
	"github.com/siriusol/simple-im-server/model"
)

func Ping(client *Client, seq string, message []byte) (code common.Error, data interface{}) {
	code = common.OK
	fmt.Println("websocket request ping", client.Addr, seq, message)
	data = "pong"
	return
}

func Login(client *Client, seq string, message []byte) (code common.Error, data interface{}) {
	code = common.OK
	// currentTS := time.Now().Unix()

	request := &model.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParamIllegal
		fmt.Println("用户登录, 解析数据失败", seq, err)
		return
	}

	fmt.Println("websocket request 用户登录", seq, "ServiceToken", request.ServiceToken)
	return nil, nil
}

func Heartbeat(client *Client, seq string, message []byte) (code common.Error, data interface{}) {
	return nil, nil
}

func Register(string, func(*Client, string, []byte) (common.Error, interface{})) {

}
