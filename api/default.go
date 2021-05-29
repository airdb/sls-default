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

// @Security ApiKeyAuth
// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func DefaultString(c *gin.Context) {
	modules := []string{
		"airdb/scf-noah",
		"airdb/scf-airdb",
		"airdb/scf-oauth2",
		"airdb/scf-wechat",
		"airdb/scf-mina",
		"airdb/scf-svgo",
	}

	c.String(http.StatusOK, strings.Join(modules, "\n"))
}
