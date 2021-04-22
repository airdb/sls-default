package api

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-airdb/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ListDailyEnglish(c *gin.Context) {
	var req vo.ListDailyEnglishReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.ListDailyEnglish(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func CreateDailyEnglish(c *gin.Context) {
	var req vo.CreateDailyEnglishReq

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	resp := vo.CreateDailyEnglish(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}
