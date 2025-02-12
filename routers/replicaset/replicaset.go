package replicaset

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/replicaset"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/replicaset")
	Get(authGroup)
	List(authGroup)

}

func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", replicaset.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", replicaset.List)
}
