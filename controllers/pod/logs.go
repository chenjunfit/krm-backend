package pod

import (
	"bufio"
	"context"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"net/http"
	"time"
)

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

func Log(r *gin.Context) {
	info := models.Infor{}
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	clientSet, err := kubeutils.NewClientSet(kubeconfig, 0)
	if err != nil {
		info.ReturnData.Status = 400
		info.ReturnData.Message = "创建ClientSet失败: " + err.Error()
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	if info.TailLines == 0 {
		info.TailLines = 100
	}
	getOptions := corev1.PodLogOptions{
		Container: info.Container,
		Follow:    true,
		TailLines: &info.TailLines,
	}
	logRequest := clientSet.CoreV1().Pods(info.NameSpace).GetLogs(info.Name, &getOptions)
	timeoutContext, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	stream, err := logRequest.Stream(timeoutContext)
	defer stream.Close()
	if err != nil {
		info.ReturnData.Status = 400
		info.ReturnData.Message = "获取pod日志失败" + err.Error()
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	conn, err := upgrader(r.Request).Upgrade(r.Writer, r.Request, nil)
	defer conn.Close()
	if err != nil {
		info.ReturnData.Status = 400
		info.ReturnData.Message = "WebSocket协议升级失败"
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		msg := scanner.Text()
		err = conn.WriteMessage(websocket.TextMessage, []byte(string(msg)))
		if err != nil {
			logs.Error(map[string]interface{}{"err:": err.Error()}, "写入日志数据失败")
			break
		}
	}
}
