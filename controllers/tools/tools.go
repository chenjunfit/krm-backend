package tools

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"net/http"
	"time"
)

func Yaml(r *gin.Context) {
	info := models.Infor{}
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	instance, err := kubeutils.NewTools(kubeconfig)
	if err != nil {
		logs.Error(map[string]interface{}{"err:": err.Error()}, "tools初始化失败")
		info.ReturnData.Status = 400
		info.ReturnData.Message = err.Error()
		r.JSON(http.StatusOK, info.ReturnData)
	}
	var errMsg string
	switch info.Method {
	case "Create":
		{
			errMsg, err = instance.Create(info.Yaml)
		}
	case "Apply":
		{
			errMsg, err = instance.Apply(info.Yaml)
		}
	case "Update":
		{
			errMsg, err = instance.Update(info.Yaml)
		}
	}
	if errMsg != "" {
		info.ReturnData.Message = errMsg
	}
	r.JSON(http.StatusOK, info.ReturnData)

}

var upgrader = func(r *http.Request) *websocket.Upgrader {
	upgrader := &websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.WriteBufferSize = 1024
	upgrader.ReadBufferSize = 1024
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	upgrader.Subprotocols = []string{r.Header.Get("Sec-Websocket-Protocol")} //设置Sec-Websocket-Protocol
	return upgrader
}

func handleWebSocketConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, messageContent, err := conn.ReadMessage()
		if err != nil {
			break
		}
		err = conn.WriteMessage(messageType, []byte(string(messageContent)+time.Now().Format(config.TimeFormat)))
		if err != nil {
			break
		}
	}
}

func Ping(r *gin.Context) {
	conn, err := upgrader(r.Request).Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go handleWebSocketConnection(conn)

}
