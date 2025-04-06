package controllers

import (
	"errors"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"krm-backend/config"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"strconv"
	"strings"
)

func NewInfo(r *gin.Context, info *models.Infor, returnMsg string) (kubeconfig string) {
	var err error
	method := strings.ToLower(r.Request.Method)
	info.ReturnData.Message = returnMsg
	info.ReturnData.Status = 200
	if method == "get" {
		err = r.ShouldBindQuery(info)

	} else if method == "post" {
		err = r.ShouldBindJSON(info)
	} else {
		err = errors.New("不支持的方法: " + method)
	}
	logs.Info(map[string]interface{}{"msg:": info}, "列表")
	if err != nil {
		info.ReturnData.Message = "请求出错: " + err.Error()
		info.ReturnData.Status = 400
		r.JSON(200, info.ReturnData)
	}
	kubeConfigString := config.ClusterKubeConfig[info.ClusterId]

	return kubeConfigString
}

func BasicInit(r *gin.Context, item interface{}) (clientSet *kubernetes.Clientset, basicInfo *models.BasicInfo, err error) {
	info := models.BasicInfo{}
	info.Item = item
	method := strings.ToLower(r.Request.Method)
	if method == "get" {
		err = r.ShouldBindQuery(&info)
	} else if method == "post" {
		err = r.ShouldBindJSON(&info)
	} else {
		err = errors.New("不支持的方法: " + method)
	}
	if err != nil {
		return nil, basicInfo, err
	}
	if info.NameSpace == "" {
		info.NameSpace = "default"
	}
	kubeConfigString := config.ClusterKubeConfig[info.ClusterId]
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfigString))
	if err != nil {
		return nil, &info, err
	}
	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {

		return nil, &info, err
	}
	return clientSet, &info, nil
}

func CreateServiceByController(containers []coreV1.Container, labels map[string]string, namespace, name, kubeconfig, resourceType string) error {
	var service coreV1.Service
	service.Name = name
	service.Spec.Selector = make(map[string]string)
	service.Labels = make(map[string]string)
	service.Annotations = make(map[string]string)
	service.Labels = labels
	service.Annotations["kubeasy.com/autoCreateService"] = "true"
	for _, container := range containers {
		for portIndex, containerPort := range container.Ports {
			var servicePort coreV1.ServicePort
			servicePort.Port = containerPort.ContainerPort
			if containerPort.Name == "" {
				servicePort.Name = container.Name + "-" + strconv.Itoa(portIndex)
			} else {
				servicePort.Name = containerPort.Name
			}

			servicePort.Protocol = containerPort.Protocol
			service.Spec.Ports = append(service.Spec.Ports, servicePort)
		}
	}
	if resourceType == "StatefulSet" {
		service.Spec.ClusterIP = "None"
	}
	instance := kubeutils.NewService(kubeconfig, &service)
	return instance.Create(namespace)
}
