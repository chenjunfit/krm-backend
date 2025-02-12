package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"krm-backend/config"
	"krm-backend/controllers"
)

func Update(r *gin.Context) {
	returnData := config.ReturnData{}
	ns := coreV1.Namespace{}
	clientSet, _, err := controllers.BasicInit(r, &ns)
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	_, err = clientSet.CoreV1().Namespaces().Update(context.TODO(), &ns, metaV1.UpdateOptions{})
	if err != nil {
		msg := "更新ns: " + ns.Name + "失败 " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	returnData.Status = 200
	returnData.Message = "更新ns: " + ns.Name + "成功"
	r.JSON(200, returnData)
	return
}
