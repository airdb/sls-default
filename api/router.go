package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
)

var GinFaas *ginAdapter.GinFaas

func NewRouter() {
	fmt.Printf("Gin start")

	r := gin.Default()

	projectPath := "/airdb"
	r.GET(projectPath, DefaultRoot)

	API := r.Group(projectPath)
	API.GET("/", DefaultRoot)
	API.GET("/v1/hosted_zone/list", ListHostedZone)
	API.POST("/v1/hosted_zone/create", CreateHostedZone)

	GinFaas = ginAdapter.New(r)

	faas.Start(Handler)
}

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	var (
		res, _ = GinFaas.ProxyWithContext(ctx, req)
		apiRes = events.APIGatewayResponse{Body: res.Body, StatusCode: http.StatusOK}
	)

	apiRes.Headers = res.Headers

	fmt.Println(apiRes.Headers)

	// if apiRes.Headers == nil {
	// 	apiRes.Headers = make(map[string]string)
	// 	apiRes.Headers["Content-Type"] = "application/json"
	// }

	return apiRes, nil
}
