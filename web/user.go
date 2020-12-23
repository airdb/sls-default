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

type QueryUserPostReq struct {
	Address string `form:"address" json:"address"`
	Province int `form:"province" json:"province"`
}

type QueryUserPostResp struct {
	Address string
	Province int
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

func UserPost(c *gin.Context)  {
	var req QueryUserPostReq
	if err := c.ShouldBind(&req); err!=nil {
		c.JSON(http.StatusBadRequest, c.Request)

		return
	}

	resp := QueryUserPostResp{
		Address:  req.Address,
		Province: req.Province,
	}
	c.JSON(http.StatusOK, resp)
}
