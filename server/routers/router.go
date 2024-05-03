package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zilanlann/acmer-manage-system/server/docs"
	"github.com/zilanlann/acmer-manage-system/server/middleware"
	"github.com/zilanlann/acmer-manage-system/server/routers/api"
	v1 "github.com/zilanlann/acmer-manage-system/server/routers/api/v1"
	"github.com/zilanlann/acmer-manage-system/server/service/casbin_service"
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
		casbin_service.AddRouterGet(apiv1, "/test", v1.Test, "admin")
		casbin_service.AddRouterGet(apiv1, "/all-user-status", v1.AllUserStatus, "admin", "teacher", "acmer")
		casbin_service.AddRouterGet(apiv1, "/users", v1.AllUsersList, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user", v1.AddUser, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/user/:id", v1.DeleteUser, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/users", v1.DeleteUsers, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user/:id", v1.UpdateUser, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user/:id/role", v1.UpdateUserRole, "admin")
		casbin_service.AddRouterPut(apiv1, "/user/:id/password", v1.UpdatePassword, "admin", "teacher")
	}
	return r
}
