package main

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	_ "krm-backend/config"
	_ "krm-backend/controllers/initcontroller"
	"krm-backend/middlewares/auth"
	"krm-backend/routers"
)

func main() {
	r := gin.Default()
	r.Use(auth.JWTCheck)
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
