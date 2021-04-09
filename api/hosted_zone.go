package web

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-mina/model/vo"
	"github.com/gin-gonic/gin"
)

func ListHostedZone(c *gin.Context) {
	var req vo.ListHostedZoneReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.ListHostedZone(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}
