package ingressclass

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	networkV1 "k8s.io/api/networking/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func List(r *gin.Context) {
	info := models.Infor{}
	ingressClass := networkV1.IngressClass{}
	info.Item = &ingressClass
	kubeconfig := controllers.NewInfo(r, &info, "查询成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngressClass(kubeconfig, &ingressClass)
	kubeUtilser = instance
	info.List(r, kubeUtilser)
}
