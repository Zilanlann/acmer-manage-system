package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
)

func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}
	
	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"expires":      "asdadadad",
	})
}