package statefulset

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Del(r *gin.Context) {
	info := models.Infor{}
	statefulSet := appV1.StatefulSet{}
	info.Item = &statefulSet
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(r, kubeUtilser)
}
