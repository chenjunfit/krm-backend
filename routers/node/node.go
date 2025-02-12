package node

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/node"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/node")
	Update(authGroup)
	Get(authGroup)
	List(authGroup)

}
func Update(authGroup *gin.RouterGroup) {
	authGroup.POST("/update", node.Update)
}
func Get(authGroup *gin.RouterGroup) {
	authGroup.GET("/get", node.Get)
}
func List(authGroup *gin.RouterGroup) {
	authGroup.GET("/list", node.List)
}
