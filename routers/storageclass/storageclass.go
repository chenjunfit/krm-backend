package storageclass

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/storageclass"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/storageclass")
	List(authGroup)
	Get(authGroup)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", storageclass.List)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", storageclass.Get)
}
