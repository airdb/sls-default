package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/airdb/scf-mina/api"
	"github.com/airdb/scf-mina/internal/version"
	"github.com/airdb/scf-mina/model/po"
	"github.com/serverless-plus/tencent-serverless-go/events"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
)

var ginFaas *ginAdapter.GinFaas

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	var (
		res, _ = ginFaas.ProxyWithContext(ctx, req)
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

func main() {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Shanghai")

	//set timezone
	version.CreatedAt = time.Now().In(loc)
	fmt.Println(os.Getenv("AIRDB_DB_MINA_API_READ"))
	po.InitDB()

	web.NewRouter()
}
