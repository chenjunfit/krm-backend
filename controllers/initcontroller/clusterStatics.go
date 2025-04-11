package initcontroller

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"krm-backend/utils/logs"
	"time"
)

func clusterStatics(clusterId, kubeconfig string) {
	clientSet, err := kubeutils.NewClientSet(kubeconfig, 10)
	if err != nil {
		logs.Error(map[string]interface{}{"err": err.Error()}, "创建clientset失败")
		return
	}
	informerFactor := informers.NewSharedInformerFactory(clientSet, time.Second*30)
	nodeInformer := informerFactor.Core().V1().Nodes().Informer()
	nsInformer := informerFactor.Core().V1().Namespaces().Informer()
	podInformer := informerFactor.Core().V1().Pods().Informer()
	deploymentInformer := informerFactor.Apps().V1().Deployments().Informer()
	daemonSetInformer := informerFactor.Apps().V1().DaemonSets().Informer()
	ClusterStaticsMap[clusterId] = make(map[string]int)
	nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["node"] += 1
		},
		DeleteFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["node"] -= 1
		},
	})

	nsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["namespace"] += 1
		},
		DeleteFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["namespace"] -= 1
		},
	})

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["pod"] += 1
		},
		DeleteFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["pod"] -= 1
		},
	})

	deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["deployment"] += 1
		},
		DeleteFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["deployment"] -= 1
		},
	})
	daemonSetInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["daemonset"] += 1
		},
		DeleteFunc: func(obj interface{}) {
			ClusterStaticsMap[clusterId]["daemonset"] -= 1
		},
	})
	stopper := make(chan struct{})
	defer close(stopper)
	FactoryStopMap[clusterId] = stopper
	go informerFactor.Start(stopper)
	<-stopper
}
