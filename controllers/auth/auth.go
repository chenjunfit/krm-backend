package auth

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"
)

type UserInfo struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	returnData := config.NewReturnData()
	userInfo := &UserInfo{}
	if err := r.ShouldBindJSON(userInfo); err != nil {
		returnData.Message = err.Error()
		returnData.Status = 401
		r.JSON(200, returnData)
		return
	}
	if userInfo.UserName == config.AdminUserName && userInfo.Password == config.AdminPassword {
		token, err := jwtutil.GenToken(userInfo.UserName)
		if err != nil {
			returnData.Message = err.Error()
			returnData.Status = 401
			r.JSON(200, returnData)
			logs.Error(map[string]interface{}{"用户名": userInfo.UserName}, "生成token失败")
			return
		}
		data := make(map[string]interface{})
		data["token"] = token
		returnData.Message = "登录成功"
		returnData.Status = 200
		returnData.Data = data
		r.JSON(200, returnData)
		return
	} else {
		returnData.Message = "用户名或密码错误"
		returnData.Status = 401
		r.JSON(200, returnData)
		return
	}
}
func Logout(r *gin.Context) {
	r.JSON(200, gin.H{
		"message": "用户退出",
		"status":  200,
	})
	logs.Debug(nil, "退出登录信息")
	return
}
