package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

type oldToken struct {
	RefreshToken string `form:"refreshToken" json:"refreshToken" binding:"required"`
}

//	@Summary		Refresh access and refresh tokens
//	@Description	refreshes tokens based on the provided refresh token
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		string			true	"Refresh Token"
//	@Success		200				{object}	app.Response	"Returns new access and refresh tokens along with their expiry time"
//	@Failure		400				{object}	app.Response	"Invalid Parameters"
//	@Failure		10007			{object}	app.Response	"Error refreshing token"
//	@Router			/refresh-token [post]
func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}

	var o oldToken
	if err := c.ShouldBind(&o); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	aToken, rToken, exTime, err := utils.RefreshToken(o.RefreshToken)
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, e.INVALID_REFRESH_TOKEN, nil)
		return
	}
	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"expires":      exTime,
	})
}
