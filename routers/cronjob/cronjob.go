package cronjob

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/cronjob"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/cronjob")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", cronjob.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", cronjob.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", cronjob.DeleteList)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", cronjob.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", cronjob.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", cronjob.List)
}
