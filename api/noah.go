package api

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-airdb/model/vo"
	"github.com/gin-gonic/gin"
)

func ListNoah(c *gin.Context) {
	var req vo.ListLinuxShellReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.InfoList()

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func ListNoahTree(c *gin.Context) {
	var req vo.ListLinuxShellReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.ListTree()

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func QueryUser(c *gin.Context) {
	var req vo.ListLinuxShellReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	resp := vo.QueryUser()

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}
