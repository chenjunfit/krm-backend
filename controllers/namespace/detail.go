package namespace

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/controllers/cluster"
)

func Detail(r *gin.Context) {
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(r, nil)
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	_, err = clientSet.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metaV1.GetOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	podItems, _ := clientSet.CoreV1().Pods(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})
	serviceItems, _ := clientSet.CoreV1().Services(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})
	secretItems, _ := clientSet.CoreV1().Secrets(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})
	configMapItems, _ := clientSet.CoreV1().ConfigMaps(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})

	deploymentItems, _ := clientSet.AppsV1().Deployments(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})
	daemonSetItems, _ := clientSet.AppsV1().DaemonSets(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})
	statefulSetItems, _ := clientSet.AppsV1().StatefulSets(basicInfo.Name).List(context.TODO(), metaV1.ListOptions{})

	totalMap := make(map[string]int)
	totalMap["pod"] = len(podItems.Items)
	totalMap["service"] = len(serviceItems.Items)
	totalMap["secret"] = len(secretItems.Items)
	totalMap["configmap"] = len(configMapItems.Items)
	totalMap["deployment"] = len(deploymentItems.Items)
	totalMap["daemonset"] = len(daemonSetItems.Items)
	totalMap["statefulset"] = len(statefulSetItems.Items)

	nameSpaceDetailList := make([]cluster.ResourceDetail, 0)
	for k, v := range totalMap {
		nameSpaceDetailList = append(nameSpaceDetailList, cluster.ResourceDetail{
			ResourceType: k,
			Total:        v,
		})
	}
	returnData.Status = 200
	returnData.Message = "查询ns详情成功"
	returnData.Data = make(map[string]interface{})
	returnData.Data["items"] = nameSpaceDetailList
	r.JSON(200, returnData)
	fmt.Println(returnData, "返回的数据")
	return
}
