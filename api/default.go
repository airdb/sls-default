package api

import (
	"net/http"

	"github.com/airdb/scf-airdb/internal/version"
	"github.com/gin-gonic/gin"
)

func DefaultRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deploy_info": version.GetBuildInfo(),
	})
}

func DefaultString(c *gin.Context) {
	c.String(http.StatusOK, "this is a string")
}
