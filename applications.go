package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleListApplications(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-applications.html", gin.H{
		"title": "Applications",
	})
}

func handleNewApplicationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new-application.html", gin.H{"title": "New Application"})
}

func handleCreateApplication(c *gin.Context) {

	type OIDCApplication struct {
		Name         string   `form:"name" binding:"required" validate:"required,alphanumunicode"`
		AppType      string   `form:"visibility" binding:"required" validate:"oneof=public private"`
		RedirectURIs []string `form:"redirect_uri" binding:"gt=0,dive,uri"`
	}

	var form OIDCApplication
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appName": form.Name, "type": form.AppType, "redirects": form.RedirectURIs})
}
