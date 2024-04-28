package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/model"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/cf"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/cf_service"
)

func AllUserStatus(c *gin.Context) {
	appG := app.Gin{C: c}

	users, _ := model.GetAllNormalUsers()
	queryUsers := []string{}
	for _, user := range users {
		queryUsers = append(queryUsers, user.CFHandle)
	}
	cf.RefreshUserInfos(queryUsers)
	allStatus := make([]cf_service.UserStatus, 0, len(users))
	for _, user := range users {
		var status cf_service.UserStatus
		if user.CFHandle != "" {
			status.CFHandle = user.CFHandle
			status.RealName = user.RealName
			status.UserName = user.Username
		}
		status.GetUserStatus()
		allStatus = append(allStatus, status)
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"userdata": allStatus,
	})
}
