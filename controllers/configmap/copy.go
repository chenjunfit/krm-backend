package configmap

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"krm-backend/utils/logs"
)

func Copy(srcKubeconfig, desKubeconfig, srcNamespace, desNamespace string, list []string) {
	srcInstance := kubeutils.NewConfigMap(srcKubeconfig, nil)
	desInstance := kubeutils.NewConfigMap(desKubeconfig, nil)
	for _, name := range list {
		logs.Debug(map[string]interface{}{"资源类型": "ConfigMap", "数据项目:": name}, "开始拷贝数据")
		_, err := srcInstance.Get(srcNamespace, name)
		if err != nil {
			logs.Error(map[string]interface{}{"资源类型": "ConfigMap", "数据项目:": name, "命名空间": desNamespace, "msg": err.Error()}, "源数据查询失败")
			continue
		}
		desInstance.Item = srcInstance.Item
		desInstance.Item.Namespace = desNamespace
		err = desInstance.Create(desNamespace)
		if err != nil {
			logs.Error(map[string]interface{}{"资源类型": "ConfigMap", "数据项目:": name, "命名空间": desNamespace, "msg": err.Error()}, "源数据拷贝失败")
		}
	}
}
