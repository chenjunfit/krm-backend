package deployment

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
	deployment := appV1.Deployment{}
	info.Item = &deployment
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	if info.AutoCreateService {
		err := controllers.CreateServiceByController(
			deployment.Spec.Template.Spec.Containers,
			deployment.Spec.Selector.MatchLabels,
			info.NameSpace,
			deployment.Name,
			kubeconfig,
			"Deployment",
		)
		if err != nil {
			logs.Error(map[string]interface{}{"err": err.Error()}, "自动创建Service失败")
			//info.ReturnData.Message = "自动创建servie失败" + err.Error()
		}
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, &deployment)

	kubeUtilser = instance
	info.Create(r, kubeUtilser)
}
