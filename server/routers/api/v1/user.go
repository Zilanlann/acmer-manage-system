package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/user_service"
)

func GetAllUserList(c *gin.Context) {
	appG := app.Gin{C: c}

	userList := user_service.UserList{}
	userList.Get()

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":        userList.Users,
		"total":       userList.Total,
		"pageSize":    10,
		"currentPage": 1,
	})
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	var user user_service.User
	if err := c.ShouldBind(&user); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	password := struct {
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&password); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	user.Password = password.Password

	if err := user.Add(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.USER_ALREADY_EXIST, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	var user user_service.User
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user.ID = uint(id)

	if err := c.ShouldBind(&user); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := user.Update(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_UPDATE_FAIL, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var user user_service.User

	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user.ID = uint(id)

	if err := user.Delete(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_DELETE_FAIL, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	var user user_service.User

	// 将字符串类型的 id 转换为 uint 类型
	ids := struct {
		Ids []uint `json:"ids" form:"ids" binding:"required"`
	}{}
	if err := c.ShouldBind(&ids); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	for _, id := range ids.Ids {
		user.ID = id
		if err := user.Delete(); err != nil {
			global.LOG.Error(err.Error())
			appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_DELETE_FAIL, nil)
			return
		}
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdateUserRole(c *gin.Context) {
	appG := app.Gin{C: c}

	var user user_service.User
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user.ID = uint(id)

	updateRole := struct {
		Role string `json:"role" form:"role" binding:"required"`
	}{}
	if err := c.ShouldBind(&updateRole); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user.Role = updateRole.Role

	if err := user.UpdateRole(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_USER_DELETE_FAIL, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdatePassword(c *gin.Context) {
	appG := app.Gin{C: c}

	var user user_service.User
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user.ID = uint(id)

	updatePassword := struct {
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&updatePassword); err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	user.Password = updatePassword.Password

	if err := user.UpdatePassword(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.ERROR_UPDATE_PASSWORD_FAIL, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func GetAllTeacherList(c *gin.Context) {
	appG := app.Gin{C: c}

	userList := user_service.UserList{}
	if err := userList.GetTeachers(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list": userList.Users,
	})
}

func GetAllAcmerList(c *gin.Context) {
	appG := app.Gin{C: c}

	userList := user_service.UserList{}
	if err := userList.GetACMers(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list": userList.Users,
	})
}
