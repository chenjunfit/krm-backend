package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"krm-backend/config"
	"krm-backend/models"
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
