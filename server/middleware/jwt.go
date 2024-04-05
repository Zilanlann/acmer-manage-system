package middleware

import (
	"net/http"

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
