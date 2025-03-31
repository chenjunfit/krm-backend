package service

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/service"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/service")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)
	Update(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", service.Add)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", service.Update)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", service.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", service.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", service.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", service.List)
}
