package statefulset

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
	statefulSet := appV1.StatefulSet{}
	info.Item = &statefulSet
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	if info.AutoCreateService {
		err := controllers.CreateServiceByController(
			statefulSet.Spec.Template.Spec.Containers,
			statefulSet.Spec.Selector.MatchLabels,
			info.NameSpace,
			statefulSet.Name,
			kubeconfig,
			"StatefulSet",
		)
		if err != nil {
			logs.Error(map[string]interface{}{"err": err.Error()}, "自动创建Service失败")
		}
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, &statefulSet)
	kubeUtilser = instance
	info.Create(r, kubeUtilser)
}
