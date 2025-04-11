package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers/initcontroller"
)

type StaticsData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type StaticsItem struct {
	ResourceType string        `json:"resourceType"`
	Header       string        `json:"header"`
	Total        int           `json:"total"`
	Data         []StaticsData `json:"data"`
}

func gernateStatics(clusterId, resourceType, displayName string, si *StaticsItem) {
	count := initcontroller.ClusterStaticsMap[clusterId][resourceType]
	staticData := StaticsData{
		Name:  displayName,
		Value: count,
	}
	si.Total += count
	si.Data = append(si.Data, staticData)

}

func Statics(r *gin.Context) {
	clientSet := config.ClientSet
	returnData := config.ReturnData{}
	options := metaV1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	secretList, err := clientSet.CoreV1().Secrets(config.MetaNamespace).List(context.TODO(), options)
	if err != nil {
		returnData.Message = "获取集群统计信息失败: " + err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	var (
		activeTotal   int
		inactiveTotal int
	)
	staticsItemList := make([]StaticsItem, 0)
	nodeStaticsItem := StaticsItem{
		ResourceType: "node",
		Header:       "节点统计",
		Total:        0,
		Data:         []StaticsData{},
	}
	nameSpaceStaticsItem := StaticsItem{
		ResourceType: "namespace",
		Header:       "命名空间统计",
		Total:        0,
		Data:         []StaticsData{},
	}
	deploymentStaticsItem := StaticsItem{
		ResourceType: "deployment",
		Header:       "Deployment统计",
		Total:        0,
		Data:         []StaticsData{},
	}
	daemonSetStaticsItem := StaticsItem{
		ResourceType: "daemonset",
		Header:       "DaemonSet统计",
		Total:        0,
		Data:         []StaticsData{},
	}
	podStaticsItem := StaticsItem{
		ResourceType: "pod",
		Header:       "Pod统计",
		Total:        0,
		Data:         []StaticsData{},
	}
	for _, clusterConfig := range secretList.Items {
		clusterId := clusterConfig.Name
		clusterDisplayName := clusterConfig.Annotations["displayName"]
		if clusterConfig.Annotations["status"] == "Active" {
			activeTotal += 1
		} else {
			inactiveTotal += 1
		}
		gernateStatics(clusterId, "node", clusterDisplayName, &nodeStaticsItem)
		gernateStatics(clusterId, "pod", clusterDisplayName, &podStaticsItem)
		gernateStatics(clusterId, "deployment", clusterDisplayName, &deploymentStaticsItem)
		gernateStatics(clusterId, "daemonset", clusterDisplayName, &daemonSetStaticsItem)
		gernateStatics(clusterId, "namespace", clusterDisplayName, &nameSpaceStaticsItem)

	}
	clusterInActiveData := StaticsData{
		Name:  "InActive",
		Value: inactiveTotal,
	}
	clusterActiveData := StaticsData{
		Name:  "Active",
		Value: activeTotal,
	}
	clusterStaticsItem := StaticsItem{
		ResourceType: "cluster",
		Header:       "集群统计",
		Total:        len(secretList.Items),
		Data:         []StaticsData{},
	}
	clusterStaticsItem.Data = append(clusterStaticsItem.Data, clusterInActiveData, clusterActiveData)
	staticsItemList = append(staticsItemList, clusterStaticsItem, nodeStaticsItem, podStaticsItem, nameSpaceStaticsItem, deploymentStaticsItem, daemonSetStaticsItem)
	returnData.Data = make(map[string]interface{})
	returnData.Data["items"] = staticsItemList
	r.JSON(200, returnData)
	return
}
