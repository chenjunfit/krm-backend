package replicaset

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Get(r *gin.Context) {
	info := models.Infor{}
	replicaset := appV1.ReplicaSet{}
	info.Item = &replicaset
	kubeconfig := controllers.NewInfo(r, &info, "查询详情成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, &replicaset)
	kubeUtilser = instance
	info.Get(r, kubeUtilser)
}
