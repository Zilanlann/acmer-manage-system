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

type contestantReq struct {
	TeamID uint `json:"teamID"`
	UserID uint `json:"userID"`
}

// AllContestList retrieves all contests and use preload to return all corresponding data in a JSON response.
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

func CreateContest(c *gin.Context) {
	appG := app.Gin{C: c}

	var contest contest_service.Contest
	if err := c.ShouldBind(&contest); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := contest.Create(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdateContest(c *gin.Context) {
	appG := app.Gin{C: c}

	var contest contest_service.Contest
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contest.ID = uint(id)

	if err := c.ShouldBind(&contest); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := contest.Update(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteContest(c *gin.Context) {
	appG := app.Gin{C: c}

	var contest contest_service.Contest
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contest.ID = uint(id)

	if err := contest.Delete(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func CreateTeam(c *gin.Context) {
	appG := app.Gin{C: c}

	var team contest_service.Team
	newTeam := struct {
		ContestID uint   `json:"contestID" form:"contestID" binding:"required"`
		ZhName    string `json:"zhName" form:"zhName" binding:"required"`
		EnName    string `json:"enName" form:"enName"`
		CoachID   uint   `json:"coachID" form:"coachId"`
		Desc      string `json:"desc" form:"desc"`
	}{}
	if err := c.ShouldBind(&newTeam); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	team.CoachID = newTeam.CoachID
	team.ContestID = newTeam.ContestID
	team.ZhName = newTeam.ZhName
	team.EnName = newTeam.EnName
	team.Desc = newTeam.Desc

	if err := team.Create(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdateTeam(c *gin.Context) {
	appG := app.Gin{C: c}

	var team contest_service.Team
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	team.ID = uint(id)

	newTeam := struct {
		ZhName  string `json:"zhName" form:"zhName" binding:"required"`
		EnName  string `json:"enName" form:"enName"`
		CoachID uint   `json:"coachId" form:"coachId"`
		Desc    string `json:"desc" form:"desc"`
	}{}
	if err := c.ShouldBind(&newTeam); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	team.CoachID = newTeam.CoachID
	team.ZhName = newTeam.ZhName
	team.EnName = newTeam.EnName
	team.Desc = newTeam.Desc

	if err := team.Update(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteTeam(c *gin.Context) {
	appG := app.Gin{C: c}

	var team contest_service.Team
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	team.ID = uint(id)

	if err := team.Delete(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func CreateContestant(c *gin.Context) {
	appG := app.Gin{C: c}

	var contestant contest_service.Contestant
	var contestantReq contestantReq
	if err := c.ShouldBind(&contestantReq); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contestant.TeamID = contestantReq.TeamID
	contestant.UserID = contestantReq.UserID

	if err := contestant.Create(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func UpdateContestant(c *gin.Context) {
	appG := app.Gin{C: c}

	var contestant contest_service.Contestant
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contestant.ID = uint(id)

	var contestantReq contestantReq
	if err := c.ShouldBind(&contestantReq); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contestant.UserID = contestantReq.UserID

	if err := contestant.Update(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteContestant(c *gin.Context) {
	appG := app.Gin{C: c}

	var contestant contest_service.Contestant
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	contestant.ID = uint(id)

	if err := contestant.Delete(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}
