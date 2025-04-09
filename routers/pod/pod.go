package pod

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/pod"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/pod")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)
	Log(authGroup)
	Exec(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", pod.Add)
}
func Log(authGroup *gin.RouterGroup) {
	authGroup.GET("/log", pod.Log)
}
func Exec(authGroup *gin.RouterGroup) {
	authGroup.GET("/exec", pod.Exec)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", pod.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", pod.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", pod.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", pod.List)
}
