package daemonset

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/daemonset"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/daemonset")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", daemonset.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", daemonset.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", daemonset.DeleteList)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", daemonset.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", daemonset.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", daemonset.List)
}
