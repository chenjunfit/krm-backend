package storageclass

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	storageV1 "k8s.io/api/storage/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func List(r *gin.Context) {
	info := models.Infor{}
	storage := storageV1.StorageClass{}
	info.Item = &storage
	kubeconfig := controllers.NewInfo(r, &info, "查询成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, &storage)
	kubeUtilser = instance
	info.List(r, kubeUtilser)
}
