package configmap

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/configmap"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/configmap")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", configmap.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", configmap.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", configmap.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", configmap.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", configmap.List)
}
