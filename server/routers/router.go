package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zilanlann/acmer-manage-system/server/docs"
	"github.com/zilanlann/acmer-manage-system/server/middleware"
	"github.com/zilanlann/acmer-manage-system/server/model"
	"github.com/zilanlann/acmer-manage-system/server/routers/api"
	v1 "github.com/zilanlann/acmer-manage-system/server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.ZapLogger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	noAuth := r.Group("/api")
	noAuth.POST("/login", api.Login)
	noAuth.POST("/register", api.Register)
	noAuth.POST("/refresh-token", api.RefreshToken)
	noAuth.POST("/verify-email", api.SendVerifyCode)
	
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWTAuth(), middleware.CheckPermission())
	{
		model.Casbin.AddPolicy("admin", "/api/v1/test", "GET")
		apiv1.GET("/test", v1.Test)

		model.Casbin.AddPolicy("acmer", "/api/v1/all-user-status", "GET")
		model.Casbin.AddPolicy("admin", "/api/v1/all-user-status", "GET")
		apiv1.GET("/all-user-status", v1.AllUserStatus)
	}
	return r
}
