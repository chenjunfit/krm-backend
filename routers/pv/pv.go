package pv

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/pv"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/persistentvolume")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", pv.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", pv.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", pv.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", pv.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", pv.List)
}
