package casbin_service

import (
	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
)

func AddRouterGet(r *gin.RouterGroup, relativePath string, handlers gin.HandlerFunc, role ...string) {
	for _, v := range role {
		global.Casbin.AddPolicy(v, relativePath, "GET")
	}
	r.GET(relativePath, handlers)
}

func AddRouterPost(r *gin.RouterGroup, relativePath string, handlers gin.HandlerFunc, role ...string) {
	for _, v := range role {
		global.Casbin.AddPolicy(v, relativePath, "POST")
	}
	r.POST(relativePath, handlers)
}

func AddRouterPut(r *gin.RouterGroup, relativePath string, handlers gin.HandlerFunc, role ...string) {
	for _, v := range role {
		global.Casbin.AddPolicy(v, relativePath, "PUT")
	}
	r.PUT(relativePath, handlers)
}

func AddRouterDelete(r *gin.RouterGroup, relativePath string, handlers gin.HandlerFunc, role ...string) {
	for _, v := range role {
		global.Casbin.AddPolicy(v, relativePath, "DELETE")
	}
	r.DELETE(relativePath, handlers)
}