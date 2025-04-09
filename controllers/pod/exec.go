package pod

import (
	"context"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"net/http"
)

type wsMessage struct {
	MessageType string `json:"messageType" form:"messageType"`
	Data        string `json:"data" form:"data"`
	Rows        uint16 `json:"rows" form:"rows"`
	Cols        uint16 `json:"cols" form:"cols"`
}

type ptyHandler struct {
	ws         *websocket.Conn
	resizeChan chan remotecommand.TerminalSize
}

func (pty *ptyHandler) Next() *remotecommand.TerminalSize {
	size := <-pty.resizeChan
	return &size

}

func (pty *ptyHandler) Write(p []byte) (n int, err error) {
	returnDataByte := make([]byte, len(p))
	n = copy(returnDataByte, p)
	err = pty.ws.WriteMessage(websocket.TextMessage, returnDataByte)
	if err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "ws写入数据失败")
		return 0, err
	}
	return n, nil
}

func (pty *ptyHandler) Read(p []byte) (n int, err error) {
	var wsMessage wsMessage
	_, msg, err := pty.ws.ReadMessage()

	if err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "读取数据失败")
		return 0, err
	}
	if err := json.Unmarshal(msg, &wsMessage); err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "json解析失败")
		return 0, err
	}
	if wsMessage.MessageType == "resize" {
		var size remotecommand.TerminalSize
		size.Height = wsMessage.Rows
		size.Width = wsMessage.Cols
		pty.resizeChan <- size
		return 0, nil
	} else {
		wsDataByte := []byte(wsMessage.Data)
		n := copy(p, wsDataByte)
		return n, nil
	}
	return 0, nil
}

func Exec(r *gin.Context) {
	info := models.Infor{}
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	clientSet, err := kubeutils.NewClientSet(kubeconfig, 0)
	if err != nil {
		info.ReturnData.Status = 400
		info.ReturnData.Message = "创建ClientSet失败: " + err.Error()
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	if info.DefaultCommand == "" {
		info.DefaultCommand = "/bin/bash"
	}
	//创建命令请求
	execRequest := clientSet.CoreV1().RESTClient().Post().Resource("pods").
		Name(info.Name).Namespace(info.NameSpace).SubResource("exec").
		VersionedParams(
			&v1.PodExecOptions{
				Stdin:     true,
				Stdout:    true,
				Stderr:    true,
				TTY:       true,
				Container: info.Container,
				Command:   []string{info.DefaultCommand},
			}, scheme.ParameterCodec,
		)
	//创建命令执行器
	restConfig, _ := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	executor, err := remotecommand.NewSPDYExecutor(restConfig, "POST", execRequest.URL())
	if err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "创建Executor失败")
		info.ReturnData.Status = 400
		info.ReturnData.Message = "创建Executor失败" + err.Error()
		r.JSON(200, info.ReturnData)
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
	//创建长链接
	resiezeChan := make(chan remotecommand.TerminalSize)
	ptyHandler := &ptyHandler{ws: conn, resizeChan: resiezeChan}
	streamOptions := remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		Tty:               true,
		TerminalSizeQueue: ptyHandler,
	}
	err = executor.StreamWithContext(context.TODO(), streamOptions)
	if err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "创建长链接失败")
	}
}
