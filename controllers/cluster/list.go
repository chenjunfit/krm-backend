package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
)

func List(r *gin.Context) {
	clientSet := config.ClientSet
	returnData := config.ReturnData{}
	options := metaV1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	lists, err := clientSet.CoreV1().Secrets(config.MetaNamespace).List(context.TODO(), options)
	if err != nil {
		returnData.Message = "获取集群列表失败: " + err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}

	returnData.Data = make(map[string]interface{})
	items := make([]map[string]string, 0)
	for _, item := range lists.Items {
		items = append(items, item.Annotations)
	}
	returnData.Data["items"] = items
	r.JSON(200, returnData)
	return

}
