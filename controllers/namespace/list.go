package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
)

func List(r *gin.Context) {
	returnData := config.ReturnData{}
	clientSet, _, err := controllers.BasicInit(r, nil)
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	nsList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	returnData.Status = 200
	returnData.Message = "查询ns列表成功"
	returnData.Data = make(map[string]interface{})
	returnData.Data["items"] = nsList.Items
	r.JSON(200, returnData)
	return
}
