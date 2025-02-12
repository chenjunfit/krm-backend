package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
)

func Get(r *gin.Context) {
	clusterId := r.Query("clusterId")
	clientSet := config.ClientSet
	returnData := config.ReturnData{}
	secret, err := clientSet.CoreV1().Secrets(config.MetaNamespace).Get(context.TODO(), clusterId, metaV1.GetOptions{})
	if err != nil {
		msg := "获取集群详情失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	returnData.Message = "获取详情成功"
	returnData.Status = 200
	returnData.Data = make(map[string]interface{})
	clusterConfig := secret.Annotations
	clusterConfig["kubeconfig"] = string(secret.Data["kubeconfig"])
	returnData.Data["item"] = clusterConfig
	r.JSON(200, returnData)
	return
}
