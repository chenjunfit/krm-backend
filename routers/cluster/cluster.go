package cluster

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/cluster"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/cluster")
	Add(authGroup)
	Del(authGroup)
	Update(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Add(authGroup *gin.RouterGroup) {
	authGroup.POST("/add", cluster.Add)
}
func Del(authGroup *gin.RouterGroup) {
	authGroup.GET("/delete", cluster.Del)
}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", cluster.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", cluster.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", cluster.List)
}
