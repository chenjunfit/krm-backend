package daemonset

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appV1 "k8s.io/api/apps/v1"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"time"
)

func Restart(r *gin.Context) {
	info := models.Infor{}
	deployment := appV1.DaemonSet{}
	info.Item = &deployment
	kubeconfig := controllers.NewInfo(r, &info, "重启成功")
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &deployment)
	_, err := instance.Get(info.NameSpace, info.Name)
	if err != nil {
		logs.Error(map[string]interface{}{"资源类型": "DaemonSet", "msg": err.Error()}, "源数据查询失败")
		info.ReturnData.Status = 400
		info.ReturnData.Message = "未找到资源"
		r.JSON(200, info.ReturnData)
		return
	}
	if instance.Item.Spec.Template.Annotations == nil {
		instance.Item.Spec.Template.Annotations = make(map[string]string)
	}
	instance.Item.Spec.Template.Annotations["kubeeasy.com/restart"] = time.Now().Format(config.TimeFormat)
	kubeUtilser = instance
	info.Update(r, kubeUtilser)
}
