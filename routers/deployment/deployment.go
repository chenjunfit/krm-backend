package deployment

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/deployment"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/deployment")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", deployment.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", deployment.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", deployment.DeleteList)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", deployment.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", deployment.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", deployment.List)
}
