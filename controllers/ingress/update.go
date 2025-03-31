package ingress

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	networkV1 "k8s.io/api/networking/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Update(r *gin.Context) {
	info := models.Infor{}
	ingress := networkV1.Ingress{}
	info.Item = &ingress
	kubeconfig := controllers.NewInfo(r, &info, "更新成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, &ingress)
	kubeUtilser = instance
	info.Update(r, kubeUtilser)
}
