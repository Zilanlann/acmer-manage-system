package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/auth_service"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

type auth struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//	@Summary		User login
//	@Description	Authenticates a user and returns access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		auth			true	"Login Credentials"
//	@Success		200		{object}	app.Response	"Returns username, roles, accessToken, refreshToken, and token expiry"
//	@Failure		400		{object}	app.Response	"Invalid Parameters"
//	@Failure		401		{object}	app.Response	"Not Valid User"
//	@Failure		500		{object}	app.Response	"Internal Server Error or Failed to Generate Token"
//	@Router			/login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	var a auth
	if err := c.ShouldBind(&a); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Password: a.Password, Username: a.Username}

	ok, err := authService.Check()
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_CHECK_FAIL, nil)
		return
	}
	if !ok {
		appG.ErrorResponse(http.StatusUnauthorized, e.ERROR_NOT_VALID_USER, nil)
		return
	}

	aToken, rToken, err := utils.GenTokens(authService.UserId, authService.Username, authService.Role)
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_GEN_TOKEN, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"username":     authService.Username,
		"roles":        authService.Role,
		"accessToken":  aToken,
		"refreshToken": rToken,
		"expires":      "2030/10/30 00:00:00",
	})
}

//	@Summary		Register a new user
//	@Description	creates a new user with the provided credentials
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		auth					true	"User Credentials"
//	@Success		200		{object}	map[string]interface{}	"Successfully registered"
//	@Failure		400		{object}	map[string]interface{}	"Invalid parameters"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Router			/register [post]
func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var a auth
	if err := c.ShouldBind(&a); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	err := authService.Add()
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_CHECK_FAIL, nil)
		return
	}
	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}
