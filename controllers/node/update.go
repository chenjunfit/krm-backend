package node

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Update(r *gin.Context) {
	info := models.Infor{}
	node := corev1.Node{}
	info.Item = &node
	kubeconfig := controllers.NewInfo(r, &info, "修改成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewNode(kubeconfig, &node)
	kubeUtilser = instance
	info.Update(r, kubeUtilser)
}
