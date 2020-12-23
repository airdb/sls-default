package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryUserRootReq struct {
	Name string `form:"name" json:"name"`
	Sex int `form:"sex" json:"name"`
}

type QueryUserRootResp struct {
	Name string
	Sex int
}

func UserRoot(c *gin.Context) {
	var req QueryUserRootReq

	if err := c.ShouldBindQuery(&req); err !=nil{
		c.JSON(http.StatusBadRequest, c.Request)
		return
	}
	var resp  QueryUserRootResp
	resp.Name = req.Name
	resp.Sex = req.Sex

	c.JSON(http.StatusOK, resp)
}
