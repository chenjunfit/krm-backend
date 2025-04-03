package namespace

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/controllers/configmap"
	"krm-backend/controllers/cronjob"
	"krm-backend/controllers/daemonset"
	"krm-backend/controllers/deployment"
	"krm-backend/controllers/secret"
	"krm-backend/controllers/service"
	"krm-backend/models"
	"krm-backend/utils/logs"
)

func Copy(r *gin.Context) {
	var (
		ns  corev1.Namespace
		err error
	)
	info := models.Infor{}
	srcKubeconfig := controllers.NewInfo(r, &info, "复制成功")
	destKubeconfig := config.ClusterKubeConfig[info.ToClusterId]
	logs.Debug(map[string]interface{}{"copy数据内容": info}, "")

	ns.Name = info.ToNamespace
	nsInstance := kubeutils.NewNamespace(destKubeconfig, &ns)
	if info.CreateNamespace {
		err = nsInstance.Create(info.ToNamespace)
	} else {
		_, err = nsInstance.Get(info.ToNamespace, info.ToNamespace)
	}
	if err != nil {
		info.ReturnData.Status = 400
		info.ReturnData.Message = err.Error()
		r.JSON(200, info.ReturnData)
		return
	}
	for resourceType, resource := range info.ToResources {
		logs.Debug(map[string]interface{}{"resourceType": resourceType, "resource": resource}, "")
		switch resourceType {
		case "Deployment":
			{
				deployment.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}
		case "DaemonSet":
			{
				daemonset.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}
		case "Service":
			{
				service.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}
		case "ConfigMap":
			{
				configmap.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}
		case "Secret":
			{
				secret.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}
		case "CronJob":
			{
				cronjob.Copy(srcKubeconfig, destKubeconfig, info.NameSpace, info.ToNamespace, resource)

			}

		}
	}
	r.JSON(200, info.ReturnData)
}
