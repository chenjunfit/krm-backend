package service

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func DeleteList(r *gin.Context) {
	info := models.Infor{}
	service := coreV1.Service{}
	info.Item = &service
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, nil)
	kubeUtilser = instance
	info.DeleteList(r, kubeUtilser)
}
