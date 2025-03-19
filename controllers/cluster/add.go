package cluster

import (
	"github.com/gin-gonic/gin"
	"krm-backend/utils/logs"
)

func Add(r *gin.Context) {
	logs.Info(map[string]interface{}{"message:": r.Request.Body}, "集群连接测试")
	addOrUpdate(r, "create")
}
