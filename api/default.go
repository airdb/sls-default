package api

import (
	"net/http"
	"strings"

	"github.com/airdb/scf-go/internal/version"
	"github.com/gin-gonic/gin"
)

func DefaultRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deploy_info": version.GetBuildInfo(),
	})
}

func DefaultString(c *gin.Context) {
	modules := []string{
		"airdb/scf-noah",
		"airdb/scf-airdb",
		"airdb/scf-oauth2",
		"airdb/scf-wechat",
		"airdb/scf-mina",
		"airdb/scf-svgo",
	}

	c.String(http.StatusOK, strings.Join(modules, "<br>"))
}
