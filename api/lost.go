package web

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-wechat/model/vo"
	"github.com/gin-gonic/gin"
)

func ListLost(c *gin.Context) {
	var req vo.ListLostReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("list lost", err)

		return
	}

	resp := vo.ListLost(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func QueryArticle(c *gin.Context) {
	var req vo.QueryArticleReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("query article", err)

		return
	}

	resp := vo.QueryArticle(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}

func ListReunion(c *gin.Context) {
	var req vo.ListReunionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("list reunion", err)
		c.JSON(http.StatusBadRequest, err)

		return
	}

	resp := vo.ListReunion(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": resp})
}
