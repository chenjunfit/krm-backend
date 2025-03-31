package secret

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/secret"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/secret")
	Add(authGroup)
	Del(authGroup)
	DeleteList(authGroup)
	Get(authGroup)
	List(authGroup)
	Update(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", secret.Add)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", secret.Update)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.POST("/delete", secret.Del)
}
func DeleteList(authGroup *gin.RouterGroup) {
	authGroup.POST("/deletelist", secret.DeleteList)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", secret.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", secret.List)
}
