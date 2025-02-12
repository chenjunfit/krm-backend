package ingress

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	networkV1 "k8s.io/api/networking/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func List(r *gin.Context) {
	info := models.Infor{}
	ingress := networkV1.Ingress{}
	info.Item = &ingress
	kubeconfig := controllers.NewInfo(r, &info, "查询成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, &ingress)
	kubeUtilser = instance
	info.List(r, kubeUtilser)
}
