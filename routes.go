package main

import (
	"net/http"

	"github.com/verkaufer/oidc-copycat/oauth2"

	"github.com/gin-gonic/gin"
)

// registerRoutes manages registering HTTP routes with the router
func registerRoutes(r *gin.Engine, directoryRepo DirectoryReader) {
	r.GET("/", handleHealthCheck)
	r.GET("/admin", handleAdminIndex)

	// Manage OIDC Applications
	r.GET("/admin/applications", handleListApplications)
	r.POST("/admin/applications", handleCreateApplication)
	r.GET("/admin/applications/new", handleNewApplicationForm)

	// Manage Directory
	r.GET("/admin/users", handleListDirectory(directoryRepo))
	r.POST("/admin/users", handleCreateUser)
	r.GET("/admin/users/new", handleNewUserForm)

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
