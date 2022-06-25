package model

import (
	"fmt"
	"time"
)

const (
	heartbeatTimeout = 3 * 60
)

type UserOnline struct {
	AccIp         string `json:"acc_ip"`
	AccPort       string `json:"acc_port"`
	AppId         int32  `json:"app_id"`
	UserId        string `json:"user_id"`
	ClientIp      string `json:"client_ip"`
	ClientPort    string `json:"client_port"`
	LoginTime     int64  `json:"login_time"`
	HeartbeatTime int64  `json:"heartbeat_time"`
	LogOutTime    int64  `json:"log_out_time"`
	Qua           string `json:"qua"`
	DeviceInfo    string `json:"device_info"`
	IsLogOff      bool   `json:"is_log_off"`
}

func UserLogin(accIp, accPort string, appId int32, userId string, addr string, loginTime int64) *UserOnline {
	return &UserOnline{
		AccIp:         accIp,
		AccPort:       accPort,
		AppId:         appId,
		UserId:        userId,
		ClientIp:      addr,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		IsLogOff:      false,
	}
}

func (u *UserOnline) IsOnline() (online bool) {
	if u.IsLogOff {
		fmt.Println("用户是否在线 用户已下线", u.AppId, u.UserId)
		return
	}
	currentTime := time.Now().Unix()
	if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
		fmt.Println("用户是否在线 心跳超时", u.AppId, u.UserId, u.HeartbeatTime)
		return
	}
	return true
}

// UserIsLocal 用户是否在本台机器上
func (u *UserOnline) UserIsLocal(localIp, localPort string) bool {
	return u.AccIp == localIp && u.AccPort == localPort
}
