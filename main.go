package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/airdb/scf-go/web"
	"github.com/gin-gonic/gin"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
)

var ginFaas *ginAdapter.GinFaas

func newRouter() {
	fmt.Printf("Gin start")

	r := gin.Default()
	r.GET("/", web.DefaultRoot)
	r.GET("/hello", web.DefaultString)

	ginFaas = ginAdapter.New(r)
}

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	var (
		res, _ = ginFaas.ProxyWithContext(ctx, req)
		apiRes = events.APIGatewayResponse{Body: res.Body, StatusCode: http.StatusOK}
	)

	apiRes.Headers = res.Headers

	if apiRes.Headers == nil {
		apiRes.Headers = make(map[string]string)
		apiRes.Headers["Content-Type"] = "application/json"
	}

	return apiRes, nil
}

func main() {
	newRouter()
	faas.Start(Handler)
}
