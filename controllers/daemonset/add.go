package daemonset

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
)

func Add(r *gin.Context) {
	info := models.Infor{}
	daemonset := appV1.DaemonSet{}
	info.Item = &daemonset
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	if info.AutoCreateService {
		err := controllers.CreateServiceByController(
			daemonset.Spec.Template.Spec.Containers,
			daemonset.Spec.Selector.MatchLabels,
			info.NameSpace,
			daemonset.Name,
			kubeconfig,
			"DaemonSet",
		)
		if err != nil {
			logs.Error(map[string]interface{}{"err": err.Error()}, "自动创建Service失败")
		}
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &daemonset)
	kubeUtilser = instance
	info.Create(r, kubeUtilser)
}
