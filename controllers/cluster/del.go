package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
)

func Del(r *gin.Context) {
	clusterId := r.Query("clusterId")
	clientSet := config.ClientSet
	returnData := config.ReturnData{}
	err := clientSet.CoreV1().Secrets(config.MetaNamespace).Delete(context.TODO(), clusterId, metaV1.DeleteOptions{})
	if err != nil {
		msg := "删除secret失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400

	} else {
		delete(config.ClusterKubeConfig, clusterId)
		returnData.Message = "删除成功"
		returnData.Status = 200
	}
	r.JSON(200, returnData)
	return
}
