package auth

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"
)

func JWTCheck(r *gin.Context) {
	requestUrl := r.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		r.Next()
		return
	}
	returnData := config.NewReturnData()
	token := r.Request.Header.Get("Authorization")
	if token == "" {
		returnData.Status = 401
		returnData.Message = "请求未携带token"
		r.JSON(200, returnData)
		r.Abort()
		return
	} else {
		claims, err := jwtutil.ParseToken(token)
		if err != nil {
			logs.Error(map[string]interface{}{"token": token}, "token验证失败")
			returnData.Status = 401
			returnData.Message = "token验证失败"
			r.JSON(200, returnData)
			r.Abort()
			return
		}
		r.Set("claims", claims)
		r.Next()
		return
	}
}
