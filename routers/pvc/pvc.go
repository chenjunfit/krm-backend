package pvc

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/pvc"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/persistentvolumeclaim")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", pvc.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", pvc.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", pvc.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", pvc.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", pvc.List)
}
