package auth

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/auth"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	Login(authGroup)
	Logout(authGroup)

}
func Login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}
func Logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}
