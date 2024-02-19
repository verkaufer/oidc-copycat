package oauth2

import "github.com/gin-gonic/gin"

func RegisterHandlers(r *gin.Engine) {
	r.GET("/authorize", authorizeEndpoint)
}
