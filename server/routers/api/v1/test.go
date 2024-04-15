package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
)

func Test(c *gin.Context) {
	appG := app.Gin{C: c}

	fmt.Println(c.Request.Method, "  ", c.Request.RequestURI)
	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"expires": "asdadadad",
	})
}
