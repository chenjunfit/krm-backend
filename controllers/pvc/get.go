package pvc

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Get(r *gin.Context) {
	info := models.Infor{}
	pvc := coreV1.PersistentVolumeClaim{}
	info.Item = &pvc
	kubeconfig := controllers.NewInfo(r, &info, "查询详情成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, &pvc)
	kubeUtilser = instance
	info.Get(r, kubeUtilser)
}
