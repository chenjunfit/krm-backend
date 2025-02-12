package pv

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func DeleteList(r *gin.Context) {
	info := models.Infor{}
	pv := coreV1.PersistentVolume{}
	info.Item = &pv
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, nil)
	kubeUtilser = instance
	info.DeleteList(r, kubeUtilser)
}
