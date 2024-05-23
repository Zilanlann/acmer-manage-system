package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/pkg/mail"
	"github.com/zilanlann/acmer-manage-system/server/pkg/redis"
	"github.com/zilanlann/acmer-manage-system/server/service/auth_service"
	"github.com/zilanlann/acmer-manage-system/server/service/cf_service"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

type auth struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type register struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Realname   string `json:"realname" form:"realname" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required"`
	VerifyCode string `json:"code" form:"code" binding:"required"`
	CfHandle   string `json:"cfHandle" form:"cfHandle"`
	AtcHandle  string `json:"atcHandle" form:"atcHandle"`
	Password   string `json:"password" form:"password" binding:"required"`
}

type verify struct {
	Email string `json:"email" form:"email" binding:"required"`
}

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
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_CHECK_FAIL, nil)
		return
	}
	if !ok {
		appG.ErrorResponse(http.StatusUnauthorized, e.ERROR_NOT_VALID_USER, nil)
		return
	}

	aToken, rToken, exTime, err := utils.GenTokens(int(authService.UserId), authService.Username, authService.Role)
	if err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_GEN_TOKEN, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"id":           authService.UserId,
		"username":     authService.Username,
		"avatar":       authService.Avatar,
		"roles":        []string{authService.Role},
		"accessToken":  aToken,
		"refreshToken": rToken,
		"expires":      exTime,
	})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	var a register
	if err := c.ShouldBind(&a); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	key := fmt.Sprintf("verify-code:%s", a.Email)
	if global.REDIS.Exists(redis.Ctx, key).Val() == 0 {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_VERIFY_CODE, nil)
		return
	}
	trueCode := global.REDIS.Get(redis.Ctx, key).Val()
	if a.VerifyCode != trueCode {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_VERIFY_CODE, nil)
		return
	}
	authService := auth_service.Auth{
		Username: a.Username, Password: a.Password,
		Realname: a.Realname, Email: a.Email, CfHandle: a.CfHandle, AtcHandle: a.AtcHandle,
	}
	if authService.CfHandle != "" {
		authService.Avatar = cf_service.GetAvatar(a.CfHandle)
	} else {
		authService.Avatar = "https://userpic.codeforces.org/no-avatar.jpg"
	}
	err := authService.Add()
	if err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.USER_ALREADY_EXIST, nil)
		return
	}
	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func SendVerifyCode(c *gin.Context) {
	appG := app.Gin{C: c}

	var a verify
	if err := c.ShouldBind(&a); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	key := fmt.Sprintf("verify-code:%s", a.Email)
	if global.REDIS.Exists(redis.Ctx, key).Val() != 0 {
		verifyCode := global.REDIS.Get(redis.Ctx, key).Val()
		mail.SendCode(verifyCode, a.Email)
	} else {
		source := rand.NewSource(time.Now().UnixNano())
		localRand := rand.New(source)
		verifyCode := strconv.Itoa(localRand.Intn(900000) + 100000)
		global.REDIS.Set(redis.Ctx, key, verifyCode, 5*time.Minute)
		mail.SendCode(verifyCode, a.Email)
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}
