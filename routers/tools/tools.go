package tools

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/tools"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/tools")
	Yaml(authGroup)
	Ping(authGroup)
}
func Yaml(authGroup *gin.RouterGroup) {
	authGroup.POST("/yaml", tools.Yaml)
}
func Ping(authGroup *gin.RouterGroup) {
	authGroup.GET("/ping", tools.Ping)
}
