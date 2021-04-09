package web

import (
	"fmt"
	"net/http"

	"github.com/airdb/scf-wechat/model/vo"
	"github.com/gin-gonic/gin"
)

// Usage ...
var Usage = `bbhj 机器人使用帮助

                示例1：bbhj 4407
                示例2：bbhj 山西 4407
                示例3：bbhj 山西 运城 张

                说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。
`

func QueryRobot(c *gin.Context) {
	var req vo.QueryRobotReq
	if err := c.ShouldBindQuery(&req); err != nil {
		// c.JSON(http.StatusBadRequest, err)
		fmt.Println("query robot", err)
		fmt.Println("query robot", c.Request)
		c.JSON(http.StatusOK, c.Request)

		return
	}

	if req.Message == "" {
		c.String(http.StatusOK, Usage)
		return
	}

	msg := ""

	losts := vo.SearchLost(req.Message)

	for _, lost := range losts {
		fmt.Println(lost)
		msg += lost.Subject + "\n" + lost.DataFrom + "\n"
	}

	c.String(http.StatusOK, msg)
}
