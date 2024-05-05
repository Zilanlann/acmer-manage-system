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

		// API about user
		casbin_service.AddRouterGet(apiv1, "/all-user-status", v1.AllUserStatus, "admin", "teacher", "acmer")
		casbin_service.AddRouterGet(apiv1, "/users", v1.GetAllUserList, "admin", "teacher")
		casbin_service.AddRouterGet(apiv1, "/teachers", v1.GetAllTeacherList, "admin", "teacher")
		casbin_service.AddRouterGet(apiv1, "/acmers", v1.GetAllAcmerList, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user", v1.CreateUser, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/user/:id", v1.DeleteUser, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/users", v1.DeleteUsers, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user/:id", v1.UpdateUser, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/user/:id/role", v1.UpdateUserRole, "admin")
		casbin_service.AddRouterPut(apiv1, "/user/:id/password", v1.UpdatePassword, "admin", "teacher")

		// API about contest
		casbin_service.AddRouterGet(apiv1, "/contests", v1.AllContestList, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/contest", v1.CreateContest, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/contest/:id", v1.UpdateContest, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/contest/:id", v1.DeleteContest, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/team", v1.CreateTeam, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/team/:id", v1.UpdateTeam, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/team/:id", v1.DeleteTeam, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/contestant", v1.CreateContestant, "admin", "teacher")
		casbin_service.AddRouterPut(apiv1, "/contestant/:id", v1.UpdateContestant, "admin", "teacher")
		casbin_service.AddRouterDelete(apiv1, "/contestant/:id", v1.DeleteContestant, "admin", "teacher")
	}
	return r
}
