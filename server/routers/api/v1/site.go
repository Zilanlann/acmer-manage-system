package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/app"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
	"github.com/zilanlann/acmer-manage-system/server/service/site_service"
)

func GetSiteList(c *gin.Context) {
	appG := app.Gin{C: c}

	siteList := site_service.SiteList{}
	if err := siteList.Get(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":        siteList.Sites,
		"total":       siteList.Total,
		"pageSize":    10,
		"currentPage": 1,
	})
}

func GetSiteTypeList(c *gin.Context) {
	appG := app.Gin{C: c}

	siteTypeList := site_service.SiteTypeList{}
	if err := siteTypeList.Get(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":  siteTypeList.SiteTypes,
		"total": siteTypeList.Total,
	})
}

func CreateSite(c *gin.Context) {
	appG := app.Gin{C: c}

	var site site_service.Site
	siteInfo := struct {
		Host       string `json:"host"`
		SiteTypeID uint   `json:"siteTypeID"`
		Name       string `json:"name"`
		Desc       string `json:"desc"`
	}{}
	if err := c.ShouldBind(&siteInfo); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	site.Host = siteInfo.Host
	site.SiteTypeID = siteInfo.SiteTypeID
	site.Name = siteInfo.Name
	site.Desc = siteInfo.Desc

	if err := site.Create(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func CreateSiteType(c *gin.Context) {
	appG := app.Gin{C: c}

	var siteType site_service.SiteType
	if err := c.ShouldBind(&siteType); err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := siteType.Create(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}

func DeleteSite(c *gin.Context) {
	appG := app.Gin{C: c}

	var site site_service.Site
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		global.LOG.Warn(err.Error())
		appG.ErrorResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	site.ID = uint(id)

	if err := site.Delete(); err != nil {
		global.LOG.Error(err.Error())
		appG.ErrorResponse(http.StatusInternalServerError, e.SERVER_ERROR, nil)
	}

	appG.SuccessResponse(http.StatusOK, e.SUCCESS, nil)
}
