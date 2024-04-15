package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		atoken := c.Request.Header.Get("Authorization")
		if atoken == "" {
			appG.ErrorResponse(http.StatusBadRequest, e.ERROR_TOKEN_INVALID, nil)
			c.Abort()
			return
		}
		// 检查token是否以"Bearer "开头，如果是则去除"Bearer "前缀
		if strings.HasPrefix(atoken, "Bearer ") {
			atoken = strings.TrimPrefix(atoken, "Bearer ")
		} else {
			// 如果没有Bearer前缀，返回错误响应
			appG.ErrorResponse(http.StatusBadRequest, e.AUTH_TOKEN_REQUIRED, nil)
			c.Abort()
			return
		}
		claim, err := utils.ParseToken(atoken) // 解析 access_token
		if err != nil {
			appG.ErrorResponse(http.StatusBadRequest, e.ERROR_TOKEN_INVALID, nil)
			c.Abort()
			return
		} else {
			c.Set("role", claim.Role)
			c.Set("user_id", claim.UserID)
			c.Set("username", claim.Username)
			c.Next()
			return
		}
	}
}
