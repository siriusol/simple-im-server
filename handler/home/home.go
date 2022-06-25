package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	appIdStr := c.Query("appId")
	appIdInt64, _ := strconv.ParseInt(appIdStr, 10, 64)
	appId := int32(appIdInt64)
	if !websocket.InAppIds(appId) {
		appId = websocket.GetDefaultAppId()
	}

	fmt.Println("http_request 聊天首页", appId)

	data := gin.H{
		"title":        "聊天首页",
		"appId":        appId,
		"httpUrl":      viper.GetString("app.httpUrl"),
		"webSocketUrl": viper.GetString("app.webSocketUrl"),
	}
	c.HTML(http.StatusOK, "index.tpl", data)
}
