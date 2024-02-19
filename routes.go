package main

import (
	"net/http"
	"oidc_copycat/oauth2"

	"github.com/gin-gonic/gin"
)

// registerRoutes manages registering HTTP routes with the router
func registerRoutes(r *gin.Engine) {
	r.GET("/", handleHealthCheck)

	oauth2.RegisterHandlers(r)
}

func handleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Health OK")
}
