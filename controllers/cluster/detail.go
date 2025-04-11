package cluster

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/controllers/initcontroller"
	"net/http"
)

type ResourceDetail struct {
	ResourceType string `json:"resourceType"`
	Total        int    `json:"total"`
}

func Detail(r *gin.Context) {
	clusterId := r.Query("clusterId")
	returnData := config.ReturnData{}
	clusterDetailList := make([]ResourceDetail, 0)
	for k, v := range initcontroller.ClusterStaticsMap[clusterId] {
		clusterDetailList = append(clusterDetailList, ResourceDetail{
			ResourceType: k,
			Total:        v,
		})
	}
	returnData.Status = http.StatusOK
	returnData.Message = "查询成功"
	returnData.Data = make(map[string]interface{})
	returnData.Data["items"] = clusterDetailList
	r.JSON(http.StatusOK, returnData)
	return
}
