package api

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-airdb/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ListLinuxShell(c *gin.Context) {
	var req vo.ListLinuxShellReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.ListLinuxShell(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func CreateLinuxShell(c *gin.Context) {
	var req vo.CreateLinuxShellReq

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	resp := vo.CreateLinuxShell(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}
