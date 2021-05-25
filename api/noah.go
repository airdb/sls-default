package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListNoah(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "xx"})
}
