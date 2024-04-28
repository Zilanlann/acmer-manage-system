package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() func(c *gin.Context) {
  // CORS for https://foo.com and https://github.com origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // - Credentials share
  // - Preflight requests cached for 12 hours
  return cors.New(cors.Config{
    AllowOrigins:     []string{"https://acm.ycitoj.top", "http://localhost:8848", "http://localhost:4173"},
    AllowMethods:     []string{"PUT", "GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"origin", "content-type", "authorization", "x-requested-with"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    // AllowOriginFunc: func(origin string) bool {
    //   return origin == "https://github.com"
    // },
    MaxAge: 12 * time.Hour,
  })
}