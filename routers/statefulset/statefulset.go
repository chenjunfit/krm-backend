package statefulset

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/statefulset"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/statefulset")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)
	Restart(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", statefulset.Add)
}
func Restart(authGroup *gin.RouterGroup) {
	authGroup.POST("/restart", statefulset.Restart)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", statefulset.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", statefulset.DeleteList)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", statefulset.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", statefulset.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", statefulset.List)
}
