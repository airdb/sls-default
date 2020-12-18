package web

import (
	"net/http"

	"github.com/airdb/scf-go/internal/version"
	"github.com/gin-gonic/gin"
)

func DefaultRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deploy_info": version.GetBuildInfo(),
		"message":     "Hello Serverless Gin.",
		"query":       c.Query("q"),
	})
}

func DefaultString(c *gin.Context) {
	c.String(http.StatusOK, "this is a string")
}
