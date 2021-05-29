package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/airdb/scf-go/docs"
)

var GinFaas *ginAdapter.GinFaas

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host petstore.swagger.io
// @BasePath /v2
func NewRouter() {
	fmt.Printf("Gin start")

	r := gin.Default()

	projectPath := "/index"
	r.GET("/", DefaultRoot)
	r.GET(projectPath, DefaultString)

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

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
