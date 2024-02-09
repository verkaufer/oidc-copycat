package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerRoutes manages registering HTTP routes with the router
func registerRoutes(r *gin.Engine) {
	r.GET("/", handleHealthCheck)
}

func handleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Health OK")
}
