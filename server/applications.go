package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	oidc_copycat "github.com/verkaufer/oidc-copycat"
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

	var form oidc_copycat.OIDCApplication
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appName": form.Name, "type": form.AppType, "redirects": form.RedirectURIs})
}
