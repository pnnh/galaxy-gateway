package auth

import (
	"github.com/pnnh/quantum-go/utils/auth"
 

	"github.com/gin-gonic/gin"
)

func GetAuth(gctx *gin.Context) (string, error) {
	cookies := gctx.Request.Cookies()
	jwtToken := ""
	for _, v := range cookies {
		if v.Name == "c" {
			jwtToken = v.Value
			break
		}
	}
	// 未登录
	if len(jwtToken) < 1 {
		return "", nil
	}
	// todo 这里应该从配置里取
	return auth.ParseToken(jwtToken, "JWTKey")
}
