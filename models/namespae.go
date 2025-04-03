package models

import (
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"krm-backend/config"
)

type BasicInfo struct {
	ClusterId  string      `json:"clusterId" form:"clusterId"`
	NameSpace  string      `json:"nameSpace" form:"nameSpace"`
	Name       string      `json:"name" form:"name"`
	Item       interface{} `json:"item"`
	DeleteList []string    `json:"deleteList"`
}
type Infor struct {
	BasicInfo
	ReturnData    config.ReturnData `json:"returnData"`
	LabelSelector string            `json:"labelSelector"`
	FieldSelector string            `json:"fieldSelector"`
	ForceDelete   bool              `json:"forceDelete"`

	//namespace copy field
	ToClusterId     string              `json:"toClusterId" form:"toClusterId"`
	ToNamespace     string              `json:"toNamespace" form:"toNamespace"`
	CreateNamespace bool                `json:"createNamespace" form:"createNamespace"`
	ToResources     map[string][]string `json:"toResources" form:"toResources"`
}

func (infor *Infor) Create(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	err := utilsinterface.Create(infor.NameSpace)
	if err != nil {
		infor.ReturnData.Message = "创建失败: " + err.Error()
		infor.ReturnData.Status = 400
	}
	r.JSON(200, infor.ReturnData)
}
func (infor *Infor) Update(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	err := utilsinterface.Update(infor.NameSpace)
	if err != nil {
		infor.ReturnData.Message = "更新失败: " + err.Error()
		infor.ReturnData.Status = 400
	}
	r.JSON(200, infor.ReturnData)
}
func (infor *Infor) List(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	items, err := utilsinterface.List(infor.NameSpace, infor.LabelSelector, infor.FieldSelector)
	if err != nil {
		infor.ReturnData.Message = "查询失败: " + err.Error()
		infor.ReturnData.Status = 400
	} else {
		infor.ReturnData.Data = make(map[string]interface{})
		infor.ReturnData.Data["items"] = items
	}
	r.JSON(200, infor.ReturnData)
}
func (infor *Infor) Get(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	item, err := utilsinterface.Get(infor.NameSpace, infor.Name)
	if err != nil {
		infor.ReturnData.Message = "创建失败: " + err.Error()
		infor.ReturnData.Status = 400
	} else {
		infor.ReturnData.Data = make(map[string]interface{})
		infor.ReturnData.Data["item"] = item
	}
	r.JSON(200, infor.ReturnData)
}
func (infor *Infor) Delete(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	var gracePeriodSeconds int64
	if infor.ForceDelete {
		var s int64 = 0
		gracePeriodSeconds = s
	}

	err := utilsinterface.Delete(infor.NameSpace, infor.Name, &gracePeriodSeconds)
	if err != nil {
		infor.ReturnData.Message = "删除失败: " + err.Error()
		infor.ReturnData.Status = 400
	}
	r.JSON(200, infor.ReturnData)
}
func (infor *Infor) DeleteList(r *gin.Context, utilsinterface kubeutils.KubeUtilser) {
	var gracePeriodSeconds int64
	if infor.ForceDelete {
		var s int64 = 0
		gracePeriodSeconds = s
	}
	err := utilsinterface.DeleteList(infor.NameSpace, infor.BasicInfo.DeleteList, &gracePeriodSeconds)
	if err != nil {
		infor.ReturnData.Message = "删除失败: " + err.Error()
		infor.ReturnData.Status = 400
	}
	r.JSON(200, infor.ReturnData)
}
