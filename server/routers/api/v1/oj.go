package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/oj_service"
)

// func AllUserStatus(c *gin.Context) {
// 	appG := app.Gin{C: c}

// 	users, _ := model.GetACMersList()
// 	queryUsers := []string{}
// 	for _, user := range users {
// 		queryUsers = append(queryUsers, user.CFHandle)
// 	}
// 	cf.RefreshUserInfos(queryUsers)
// 	allStatus := make([]cf_service.UserStatus, 0, len(users))
// 	for _, user := range users {
// 		var status cf_service.UserStatus
// 		if user.CFHandle != "" {
// 			status.CFHandle = user.CFHandle
// 			status.RealName = user.RealName
// 			status.UserName = user.Username
// 		}
// 		status.GetUserStatus()
// 		allStatus = append(allStatus, status)
// 	}

// 	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
// 		"userdata": allStatus,
// 	})
// }

func AllOJContestList(c *gin.Context) {
	appG := app.Gin{C: c}

	contestList := oj_service.OJContestList{}
	if err := contestList.GetAllContests(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":  contestList.Contests,
		"total": contestList.Total,
	})
}

func GetUserSubmissionsList(c *gin.Context) {
	appG := app.Gin{C: c}

	userSubmissionsList := oj_service.OJSubmissionList{}

	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := userSubmissionsList.GetByUser(uint(id)); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":  userSubmissionsList.Submissions,
		"total": userSubmissionsList.Total,
	})
}

func GetAllTagList(c *gin.Context) {
	appG := app.Gin{C: c}

	tagList := oj_service.ProblemTagList{}

	if err := tagList.Get(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":  tagList.Tags,
		"total": tagList.Total,
	})
}
