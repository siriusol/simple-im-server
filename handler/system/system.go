package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/siriusol/simple-im-server/common"
	"runtime"
)

// Status 查询系统状态
func Status(c *gin.Context) {
	isDebug := c.Query("isDebug")
	fmt.Println("http_request 查询系统状态", isDebug)

	data := make(map[string]interface{})

	numGoroutine := runtime.NumGoroutine()
	numCPU := runtime.NumCPU()

	data["numGoroutine"] = numGoroutine
	data["numCPU"] = numCPU

	data["managerInfo"] = websocket.GetManagerInfo(isDebug)

	controllers.Response(c, common.OK, data)
}
