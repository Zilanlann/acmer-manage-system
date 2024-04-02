package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/auth_service"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username" form:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password" form:"password"`
}

// @Summary Login
// @Produce json
// @Param auth body auth true "用户登录信息"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [Post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var a auth
	if err := c.ShouldBindJSON(&a); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	ok, _ := valid.Valid(&a)

	username := a.Username
	password := a.Password

	if !ok {
		fmt.Println("valid.Errors")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Password: password, Username: username}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_CHECK_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_NOT_VALID_USER, nil)
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"roles":    authService.Role,
		"userId":   authService.UserId,
		"username": authService.Username,
		"realName": authService.RealName,
		"desc":     authService.Desc,
		"token":    token,
	})
}

// @Summary Register
// @Produce json
// @Param username formData string true "userName"
// @Param password formData string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /register [Post]
func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		fmt.Println("valid.Errors")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	err := authService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_CHECK_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
