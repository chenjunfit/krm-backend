package namespace

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/namespace"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/namespace")
	Add(authGroup)
	Del(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)
	Copy(authGroup)

}
func Copy(authGroup *gin.RouterGroup) {
	authGroup.POST("/copy", namespace.Copy)
}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", namespace.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.GET("/delete", namespace.Del)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", namespace.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", namespace.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", namespace.List)
}
