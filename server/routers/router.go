package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zilanlann/acmer-manage-system/server/docs"
	"github.com/zilanlann/acmer-manage-system/server/middleware"
	"github.com/zilanlann/acmer-manage-system/server/routers/api"
	v1 "github.com/zilanlann/acmer-manage-system/server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.POST("/refresh-token", api.RefreshToken)
	
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWTAuth())
	{
		apiv1.GET("/test", v1.Test)
	}
	return r
}
