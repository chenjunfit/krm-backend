package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
)

func Add(r *gin.Context) {
	returnData := config.ReturnData{}

	clientSet, basicInfo, err := controllers.BasicInit(r, nil)
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	ns := coreV1.Namespace{}
	ns.Name = basicInfo.Name
	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &ns, metaV1.CreateOptions{})
	if err != nil {
		msg := "创建ns: " + ns.Name + "失败 " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	returnData.Status = 200
	returnData.Message = "创建ns: " + ns.Name + "成功"
	r.JSON(200, returnData)
	return
}
