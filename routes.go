package main

import (
	"net/http"
	"oidc_copycat/oauth2"

	"github.com/gin-gonic/gin"
)

// registerRoutes manages registering HTTP routes with the router
func registerRoutes(r *gin.Engine) {
	r.GET("/", handleHealthCheck)
	r.GET("/admin", handleAdminIndex)

	oauth2.RegisterHandlers(r)
}

func handleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Health OK")
}

func handleAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
		"title": "Admin Panel",
	})
}
