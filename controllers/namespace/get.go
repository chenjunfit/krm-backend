package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
)

func Get(r *gin.Context) {
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(r, nil)
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	ns, err := clientSet.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metaV1.GetOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	returnData.Status = 200
	returnData.Message = "查询ns详情成功"
	returnData.Data = make(map[string]interface{})
	returnData.Data["item"] = ns
	r.JSON(200, returnData)
	return
}
