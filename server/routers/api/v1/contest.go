package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/contest_service"
)

func AllContestList(c *gin.Context) {
	appG := app.Gin{C: c}

	contestList := contest_service.ContestList{}
	if err := contestList.Get(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":        contestList.Contests,
		"total":       contestList.Total,
		"pageSize":    10,
		"currentPage": 1,
	})
}

func GetContestInfo(c *gin.Context) {
	appG := app.Gin{C: c}

	var contest contest_service.Contest
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contest.ID = uint(id)

	if err := contest.Get(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"contest": contest,
	})
}
