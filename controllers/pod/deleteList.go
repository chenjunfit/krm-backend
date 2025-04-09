package pod

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
)

func DeleteList(r *gin.Context) {
	info := models.Infor{}
	pod := coreV1.Pod{}
	info.Item = &pod
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	logs.Debug(map[string]interface{}{"删除测试": info}, "msg")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	kubeUtilser = instance
	info.DeleteList(r, kubeUtilser)
}
