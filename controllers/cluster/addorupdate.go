package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/utils"
	"krm-backend/utils/logs"
)

func addOrUpdate(r *gin.Context, method string) {
	args := ""
	if method == "create" {
		args = "创建集群"
	} else {
		args = "修改集群"
	}
	cluster := &ClusterConfig{}
	returnData := config.ReturnData{}
	if err := r.ShouldBindJSON(cluster); err != nil {
		msg := "集群数据不完整 " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	clusterStatus, err := cluster.GetClusterStatus()
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error(), "cluseter_id": cluster.Id}, "获取集群状态失败")
		msg := "集群连接不上: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}

	var secret coreV1.Secret
	secret.Name = cluster.Id
	secret.Labels = make(map[string]string)
	secret.Labels[config.ClusterConfigSecretLabelKey] = config.ClusterConfigSecretLabelValue

	m := utils.Struct2map(clusterStatus)
	secret.Annotations = make(map[string]string)
	secret.Annotations = m

	secret.StringData = make(map[string]string)
	secret.StringData["kubeconfig"] = cluster.KubeConfig

	clientset := config.ClientSet
	if method == "create" {
		_, err = clientset.CoreV1().Secrets(config.MetaNamespace).Create(context.TODO(), &secret, metaV1.CreateOptions{})

	} else {
		_, err = clientset.CoreV1().Secrets(config.MetaNamespace).Update(context.TODO(), &secret, metaV1.UpdateOptions{})

	}
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error(), "cluseter_id": cluster.Id, "displayname": cluster.DisplayName}, "创建secret失败")
		msg := args + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	config.ClusterKubeConfig[cluster.Id] = cluster.KubeConfig
	returnData.Message = args + "成功"
	returnData.Status = 200
	r.JSON(200, returnData)
	return

}
