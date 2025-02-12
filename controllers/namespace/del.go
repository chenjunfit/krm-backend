package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
)

func Del(r *gin.Context) {
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(r, nil)
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	for _, item := range config.ProtectNameSpace {
		if item == basicInfo.Name {
			returnData.Message = "删除被保护的ns: " + basicInfo.Name + "失败"
			returnData.Status = 400
			r.JSON(200, returnData)
			return
		}
	}
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metaV1.DeleteOptions{})
	if err != nil {
		returnData.Message = "删除ns: " + basicInfo.Name + "失败" + err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	returnData.Message = "删除ns: " + basicInfo.Name + "成功"
	returnData.Status = 200
	r.JSON(200, returnData)
	return
}
