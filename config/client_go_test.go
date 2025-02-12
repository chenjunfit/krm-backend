package config

import (
	"context"
	"encoding/json"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"krm-backend/utils/logs"
	"testing"
)
import "k8s.io/client-go/tools/clientcmd"

func TestClientGoList(t *testing.T) {
	//初始化实例
	config, err := clientcmd.BuildConfigFromFlags("", "./meta/kubeconfig.yml")
	if err != nil {
		panic(err)
	}
	//创建client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//操作集群
	pods, err := clientSet.CoreV1().Pods("kube-system").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logs.Error(nil, "查询pod失败")
	} else {
		fmt.Println("pod数量", len(pods.Items))
	}
	deployments, err := clientSet.AppsV1().Deployments("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logs.Error(nil, "查询deployment失败")
	} else {
		fmt.Println(len(deployments.Items))
		for _, deploy := range deployments.Items {
			fmt.Println(deploy.Namespace, deploy.Name)
		}
	}

	ns, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logs.Error(nil, "查询ns失败")
	} else {
		fmt.Println(len(ns.Items))
		for _, ns := range ns.Items {
			fmt.Println(ns.Name)
		}
	}

	pod, err := clientSet.CoreV1().Pods("default").Get(context.TODO(), "nginx-7854ff8877-dkbdc", metaV1.GetOptions{})
	if err != nil {
		logs.Error(nil, "查询pod失败")
	} else {
		fmt.Println(pod.Spec.Containers[0].Image)
	}

}
func TestClientGoUpdate(t *testing.T) {
	//初始化实例
	config, err := clientcmd.BuildConfigFromFlags("", "./meta/kubeconfig.yml")
	if err != nil {
		panic(err)
	}
	//创建client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deployment, err := clientSet.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metaV1.GetOptions{})
	//不能修改名字
	//deployment.Name = "newNginx"
	//_, err = clientSet.AppsV1().Deployments("default").Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//修改当前labels
	//labels := deployment.Labels
	//如果labels没有，需要初始化make()一下，不然报错空指针
	//labels["test"] = "test-update"
	//修改annotitaon
	//deployment.Annotations["test"] = "new-test"
	//修改副本数量
	//replicas := int32(2)
	//deployment.Spec.Replicas = &replicas
	//修改镜像
	deployment.Spec.Template.Spec.Containers[0].Image = "nginx:1.22.1"
	_, err = clientSet.AppsV1().Deployments("default").Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

}
func TestClientGoDel(t *testing.T) {
	//初始化实例
	config, err := clientcmd.BuildConfigFromFlags("", "./meta/kubeconfig.yml")
	if err != nil {
		panic(err)
	}
	//创建client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	err = clientSet.CoreV1().Pods("default").Delete(context.TODO(), "nginx-7854ff8877-dkbdc", metaV1.DeleteOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestClientGoCreate(t *testing.T) {
	//初始化实例
	config, err := clientcmd.BuildConfigFromFlags("", "./meta/kubeconfig.yml")
	if err != nil {
		panic(err)
	}
	//创建client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//创建namespace
	//ns := coreV1.Namespace{}
	//ns.Name = "client-go-test"

	//创建deployment
	deployment := appsV1.Deployment{}
	deployment.Name = "nginx"
	deployment.Namespace = "default"
	labels := make(map[string]string)
	labels["app"] = "nginx"
	labels["v1"] = "nginx"
	//pod选择器的label
	selector := &metaV1.LabelSelector{}
	deployment.Spec.Selector = selector
	MatchLabels := make(map[string]string)
	MatchLabels = labels
	deployment.Spec.Selector = selector
	deployment.Spec.Selector.MatchLabels = MatchLabels
	//pod的label
	deployment.Spec.Template.Labels = labels
	//depolyment的label
	deployment.Labels = labels
	//创建容器
	containers := make([]coreV1.Container, 0)
	//创建容器
	container1 := coreV1.Container{}
	container1.Image = "redis"
	container1.Name = "redis"
	containers = append(containers, container1)
	container2 := coreV1.Container{}
	container2.Image = "nginx"
	container2.Name = "nginx"
	containers = append(containers, container2)
	deployment.Spec.Template.Spec.Containers = containers
	_, err = clientSet.AppsV1().Deployments(deployment.Namespace).Create(context.TODO(), &deployment, metaV1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

}
func TestClientGoJsonCreate(t *testing.T) {
	deployJson := `{
    "kind": "Deployment",
    "apiVersion": "apps/v1",
    "metadata": {
        "name": "test",
        "creationTimestamp": null,
        "labels": {
            "app": "test"
        }
    },
    "spec": {
        "replicas": 1,
        "selector": {
            "matchLabels": {
                "app": "test"
            }
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "test"
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "nginx",
                        "image": "nginx",
                        "resources": {}
                    }
                ]
            }
        },
        "strategy": {}
    },
    "status": {}
`
	deployment := &appsV1.Deployment{}
	err := json.Unmarshal([]byte(deployJson), deployment)

	if err != nil {
		fmt.Println(err.Error())
	}
}
