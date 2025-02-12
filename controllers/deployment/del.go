package deployment

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Del(r *gin.Context) {
	info := models.Infor{}
	deployment := appV1.Deployment{}
	info.Item = &deployment
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(r, kubeUtilser)
}
