package ingress

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/ingress"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/ingress")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)
	Update(authGroup)
	Topology(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", ingress.Add)
}
func Topology(authGroup *gin.RouterGroup) {
	authGroup.GET("/topology", ingress.Topology)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", ingress.Update)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", ingress.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", ingress.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", ingress.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", ingress.List)
}
