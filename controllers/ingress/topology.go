package ingress

import (
	"fmt"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"krm-backend/controllers"
	"krm-backend/models"
	"krm-backend/utils/logs"
	"net/http"
	"strconv"
	"strings"
)

type Node struct {
	Id           string `json:"id"`
	Label        string `json:"label"`
	ResourceType string `json:"resourceType"`
}
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func Topology(r *gin.Context) {
	info := models.Infor{}
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	ingressInstance := kubeutils.NewIngress(kubeconfig, nil)
	serviceInstance := kubeutils.NewService(kubeconfig, nil)
	podInstance := kubeutils.NewPod(kubeconfig, nil)
	_, err := ingressInstance.Get(info.NameSpace, info.Name)
	if err != nil {
		logs.Warning(map[string]interface{}{"err:": err.Error(), "命名空间": info.NameSpace, "ingress": info.Name}, "获取拓扑图失败")
		info.ReturnData.Status = 400
		info.ReturnData.Message = fmt.Sprintf("获取拓图失败: %s", err.Error())
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	nodes := make([]Node, 0)
	edges := make([]Edge, 0)
	servicePod := make(map[string][]string)
	hostUsed := make(map[string]string)
	serviceUsed := make(map[string]string)
	servicePodUsed := make(map[string]string)
	podList := make([]string, 0)
	ingress := ingressInstance.Item
	//获取所有的rule规则
	rules := ingress.Spec.Rules
	for _, rule := range rules {
		host := rule.Host
		if host == "" {
			host = "*"
		}
		hostNode := Node{
			Id:           host,
			Label:        host,
			ResourceType: "domain",
		}
		if _, ok := hostUsed[host]; !ok {
			hostUsed[host] = "true"
			nodes = append(nodes, hostNode)
		}
		for _, path := range rule.HTTP.Paths {
			pathId := fmt.Sprintf("%s%s%s", host, path.Path, *path.PathType)
			pathNode := Node{
				Id:           pathId,
				Label:        path.Path,
				ResourceType: "path",
			}
			nodes = append(nodes, pathNode)
			//host到path的链接
			hostPathEdge := Edge{
				From: host,
				To:   pathId,
			}
			edges = append(edges, hostPathEdge)
			service := path.Backend.Service
			portString := ""
			if service.Port.Number == 0 {
				portString = service.Port.Name
			} else {
				portString = strconv.Itoa(int(service.Port.Number))
			}
			serviceId := fmt.Sprintf("%s:%s", service.Name, portString)

			if _, ok := serviceUsed[serviceId]; !ok {
				serviceNode := Node{
					Id:           serviceId,
					Label:        serviceId,
					ResourceType: "service",
				}
				nodes = append(nodes, serviceNode)
				serviceUsed[serviceId] = "true"
			}
			//链接path到servicce
			pathServiceEdge := Edge{
				From: pathId,
				To:   serviceId,
			}
			edges = append(edges, pathServiceEdge)
			//service pod关系处理
			if _, ok := servicePod[service.Name]; !ok {
				servicePodList := make([]string, 0)
				_, err = serviceInstance.Get(info.NameSpace, service.Name)
				if err == nil {
					selectorMap := serviceInstance.Item.Spec.Selector
					selectorList := make([]string, 0)
					for key, value := range selectorMap {
						selectorString := fmt.Sprintf("%s=%s", key, value)
						selectorList = append(selectorList, selectorString)
					}
					if len(selectorList) > 0 {
						_, err = podInstance.List(info.NameSpace, strings.Join(selectorList, ","), "")
						if err == nil {
							for _, pod := range podInstance.Items.Items {
								servicePodList = append(servicePodList, pod.Name)
								podExist := false
								for _, name := range podList {
									if pod.Name == name {
										podExist = true
										break
									}
								}
								if podExist {
									continue
								}
								podNode := Node{
									Id:           pod.Name,
									Label:        pod.Name,
									ResourceType: "pod",
								}
								nodes = append(nodes, podNode)
								podList = append(podList, pod.Name)
							}
						}
					}

				}
				servicePod[service.Name] = servicePodList
			}
			//把servie链接pod
			if _, ok := servicePodUsed[serviceId]; !ok {
				fmt.Println("serviceId: ", serviceId, service.Name, servicePod[service.Name])
				for _, podName := range servicePod[service.Name] {
					servicePodEdge := Edge{
						From: serviceId,
						To:   podName,
					}
					edges = append(edges, servicePodEdge)
					servicePodUsed[serviceId] = "true"
				}
			}

		}
	}

	info.ReturnData.Data = make(map[string]interface{})
	info.ReturnData.Data["nodes"] = nodes
	info.ReturnData.Data["edges"] = edges
	r.JSON(http.StatusOK, info.ReturnData)
	return
}
