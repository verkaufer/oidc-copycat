package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleListDirectory(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-directory.html", gin.H{
		"title": "User Directory",
	})
}

func handleNewUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new-user.html", gin.H{"title": "New User"})
}

func handleCreateUser(c *gin.Context) {
	type User struct {
		GivenName  string `form:"first_name" binding:"required" validate:"required,alphanumunicode"`
		FamilyName string `form:"last_name" binding:"required" validate:"required,alphanumunicode"`
		Email      string `form:"email" binding:"required" validate:"required,email"`
		Identifier string `form:"identifier,omitempty" validate:"ascii"`
	}

	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"givenName": form.GivenName, "familyName": form.FamilyName, "email": form.Email, "identifier": form.Identifier})
}
