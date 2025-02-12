package cluster

import (
	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	addOrUpdate(r, "create")
}
