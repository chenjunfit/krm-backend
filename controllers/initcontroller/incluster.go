package initcontroller

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"krm-backend/config"
	"krm-backend/utils/logs"
)

var ClusterStaticsMap map[string]map[string]int
var FactoryStopMap map[string]chan struct{}

func metaDataInit() {
	logs.Debug(nil, "初始化元数据")
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", "./config/meta/kubeconfig.yml")
	if err != nil {
		logs.Error(map[string]interface{}{"msg: ": err.Error()}, "incluster kubeconfig加载失败")
		panic(err.Error())
	}
	//创建client
	clientSet, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		logs.Error(map[string]interface{}{"msg: ": err.Error()}, "clientset加载失败")
		panic(err.Error())
	}
	config.ClientSet = clientSet

	var ns coreV1.Namespace
	ns.Name = config.MetaNamespace
	_, err = clientSet.CoreV1().Namespaces().Get(context.TODO(), ns.Name, metaV1.GetOptions{})
	if err != nil {
		_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &ns, metaV1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]interface{}{"msg: ": err.Error()}, "元数据namespace创建失败")
			panic(err.Error())
		}
	} else {
		inclusterversion, _ := clientSet.Discovery().ServerVersion()
		logs.Info(map[string]interface{}{"Namespace: ": config.MetaNamespace, "version": inclusterversion}, "元数据namespace已存在")
	}

	config.ClusterKubeConfig = make(map[string]string)
	options := metaV1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	lists, _ := clientSet.CoreV1().Secrets(config.MetaNamespace).List(context.TODO(), options)
	tmp := make([]string, 0)
	ClusterStaticsMap = make(map[string]map[string]int)
	FactoryStopMap = make(map[string]chan struct{})
	for _, secret := range lists.Items {
		clusterId := secret.Name
		kubeconfig := secret.Data["kubeconfig"]
		config.ClusterKubeConfig[clusterId] = string(kubeconfig)
		go clusterStatics(clusterId, string(kubeconfig))
		tmp = append(tmp, clusterId)
	}
	logs.Debug(map[string]interface{}{"当前集群配置": tmp}, "msg")

}
