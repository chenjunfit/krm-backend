package cluster

import "github.com/gin-gonic/gin"

func Update(r *gin.Context) {
	addOrUpdate(r, "update")

}
