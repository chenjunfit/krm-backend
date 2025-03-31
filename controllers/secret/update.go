package secret

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Update(r *gin.Context) {
	info := models.Infor{}
	secret := coreV1.Secret{}
	info.Item = &secret
	kubeconfig := controllers.NewInfo(r, &info, "更新成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, &secret)
	kubeUtilser = instance
	info.Update(r, kubeUtilser)
}
