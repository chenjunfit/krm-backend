package daemonset

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Add(r *gin.Context) {
	info := models.Infor{}
	daemonset := appV1.DaemonSet{}
	info.Item = &daemonset
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &daemonset)
	kubeUtilser = instance
	info.Create(r, kubeUtilser)
}
