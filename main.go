package main

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	_ "krm-backend/config"
	_ "krm-backend/controllers/initcontroller"
	"krm-backend/middlewares/auth"
	"krm-backend/middlewares/cors"
	"krm-backend/routers"
)

func main() {
	r := gin.Default()
	if config.GIN_MODE == "debug" || config.GIN_MODE == "dev" {
		r.Use(cors.Cors)
	}
	r.Use(auth.JWTCheck)
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
