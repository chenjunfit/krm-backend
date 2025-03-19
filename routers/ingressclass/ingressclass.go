package ingressclass

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/ingressclass"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/ingressclass")
	List(authGroup)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", ingressclass.List)
}
