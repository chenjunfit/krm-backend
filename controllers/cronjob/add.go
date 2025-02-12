package cronjob

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	batchV1 "k8s.io/api/batch/v1"
	"krm-backend/controllers"
	"krm-backend/models"
)

func Add(r *gin.Context) {
	info := models.Infor{}
	cronjob := batchV1.CronJob{}
	info.Item = &cronjob
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, &cronjob)
	kubeUtilser = instance
	info.Create(r, kubeUtilser)
}
