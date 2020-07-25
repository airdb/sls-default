package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ajstarks/svgo"
	"github.com/airdb/scf-go/model/vo"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

// Help Information.
var Usage = `bbhj 机器人使用帮助

		示例1：bbhj 4407
		示例2：bbhj 山西 4407
		示例3：bbhj 山西 运城 张

		说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。
`

// Refer: https://xuthus.cc/go/scf-go-runtime.html
func handler(ctx context.Context, event events.APIGatewayRequest) (interface{}, error){
	// po.InitDB()

	subPath := strings.TrimPrefix(event.Path, "/helloworld")
	switch subPath {
	case "/hello":
		fmt.Println("run hello")
		return Run(event)
	case "/robot":
		fmt.Println("run robot")
		return robot(event)
	case "", "/index":
		fmt.Println("run index")
		return defaultIndex(event)
	case "/svg":
		return svgOut(event)
	default:
		return defaultIndex(event)
	}
}

func Run(event events.APIGatewayRequest) (resp events.APIGatewayResponse, err error) {
	// 返回body中的内容
	var body = struct {
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
	}{}

	//整体API响应 保证CORS跨域
	resp = events.APIGatewayResponse{
		IsBase64Encoded: false,
		Headers:         map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "*"},
	}

	// 检测 post/get 等请求参数 tel 是否缺失 (是一个二维数组 即 key:[value...])
	// 检测 headers 参数 SecretKey 是否缺失
	// 检测 body 传递的 json数据 是否缺失
	if event.QueryString["message"] != nil {
		body.Code = 200
		body.Message = "参数齐全！"
		ret, _ := json.Marshal(body)
		resp.Body = string(ret)
		resp.StatusCode = body.Code
		return //返回
	}

	body.Code = 400
	body.Message = "error:参数缺失"
	ret, _ := json.Marshal(body)
	resp.Body = string(ret)
	resp.StatusCode = body.Code
	return
}

func defaultIndex(event events.APIGatewayRequest) (resp events.APIGatewayResponse, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html"

	resp.StatusCode = 200
	resp.Headers = headers
	resp.Body = Usage

	if event.QueryString["message"] != nil {
		return resp, nil
	}

	resp.Body ="message 参数缺失2"
	return resp, nil
}

func robot(event events.APIGatewayRequest) (resp events.APIGatewayResponse, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html"

	resp.StatusCode = 200
	resp.Headers  = headers

	if event.QueryString["message"] == nil{
		resp.Body = Usage
		return resp, nil
	}

	keywords := event.QueryString["message"]
	fmt.Println("keywords", keywords, event.QueryString["message"])

	if  len(keywords) == 1 {
		losts := vo.SearchLost(keywords[0])

		for _, lost := range losts {
			fmt.Println(lost)
			resp.Body += lost.Subject + "\n" + lost.DataFrom + "\n"
		}
	}

	return resp, nil
}

func svgOut(event events.APIGatewayRequest) (resp events.APIGatewayResponse, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "image/svg+xml"

	resp.StatusCode = 200
	resp.Headers = headers
	resp.Body = Usage

	width := 500
	height := 500

	buf := new(bytes.Buffer)
	canvas := svg.New(buf)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "About Me", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()

	resp.Body = buf.String()
	return resp, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(handler)
}
