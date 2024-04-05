package oauth2

import "github.com/gin-gonic/gin"

func RegisterHandlers(r *gin.Engine) {
	r.GET("/oauth/authorize", authorizeEndpoint)
	// r.GET("/oauth/userinfo", "")
	// r.POST("/oauth/token", "")
	// r.GET("/.well-known/openid-configuration", "")
}
