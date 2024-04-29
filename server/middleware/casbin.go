package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
)

func CheckPermission() func(c *gin.Context) {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		// 获取三个参数，角色、路由、方法
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub, _ := c.Get("role")

		// 判断策略是否存在了
		if ok, _ := global.Casbin.Enforce(sub, obj, act); ok {
			c.Next()
		} else {
			appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PERMISSION, nil)
			c.Abort()
		}
	}
}
