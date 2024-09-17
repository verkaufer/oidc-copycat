package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	oidc_copycat "github.com/verkaufer/oidc-copycat"
)

func handleNewUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new-user.html", gin.H{"title": "New User"})
}

func handleCreateUser(d *oidc_copycat.DirectoryService) gin.HandlerFunc {
	type request struct {
		GivenName  string `form:"first_name" binding:"required" validate:"required,alphanumunicode"`
		FamilyName string `form:"last_name" binding:"required" validate:"required,alphanumunicode"`
		Email      string `form:"email" binding:"required" validate:"required,email"`
		Identifier string `form:"identifier,omitempty" validate:"ascii"`
	}

	type response struct {
		User *oidc_copycat.User `json:"user"`
	}
	return func(c *gin.Context) {
		var form request
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := &oidc_copycat.User{
			FirstName:  form.GivenName,
			LastName:   form.FamilyName,
			Email:      form.Email,
			Identifier: form.Identifier,
		}
		createdUser, err := d.CreateUser(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := response{
			User: createdUser,
		}

		c.JSON(http.StatusOK, res)

	}
}

func handleListDirectory(d *oidc_copycat.DirectoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := d.ListUsers()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not ListUsers: %w", err))
			return
		}

		ctx.HTML(http.StatusOK, "admin-directory.html", gin.H{
			"title": "User Directory",
			"users": users,
		})
	}

}
