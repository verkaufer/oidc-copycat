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
	r.GET("/admin/applications", handleAdminApplications)
	r.GET("/admin/users", handleDirectoryList)

	oauth2.RegisterHandlers(r)
}

func handleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Health OK")
}

func handleAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-index.html", gin.H{
		"title": "Admin Panel",
	})
}

func handleAdminApplications(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-applications.html", gin.H{
		"title": "Applications",
	})
}

func handleDirectoryList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-directory.html", gin.H{
		"title": "User Directory",
	})
}
