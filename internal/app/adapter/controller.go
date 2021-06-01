package adapter

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/airdb/scf-go/internal/app/adapter/repository"
	"github.com/airdb/scf-go/internal/app/adapter/service"
	"github.com/airdb/scf-go/internal/app/application/usecase"
	"github.com/airdb/scf-go/internal/app/domain/valueobject"
	"github.com/airdb/scf-go/internal/version"

	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"

	"github.com/gin-gonic/gin"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	bitbank             = service.Bitbank{}
	parameterRepository = repository.Parameter{}
	orderRepository     = repository.Order{}
)

// Controller is a controller
type Controller struct{}

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

// @schemes https
// @host scf.baobeihuijia.com
// @BasePath /test/
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

/*
// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	// NOTICE: following path is from CURRENT directory, so please run Gin from root directory
	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/", ctrl.index)
	r.GET("/ticker", ctrl.ticker)
	r.GET("/candlestick", ctrl.candlestick)
	r.GET("/parameter", ctrl.parameter)
	r.GET("/order", ctrl.order)
	return r
}
*/

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Goilerplate",
	})
}

func DefaultRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deploy_info": version.GetBuildInfo(),
	})
}

// @Security ApiKeyAuth
// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /test/index [get]
func DefaultString(c *gin.Context) {
	modules := []string{
		"airdb/scf-noah",
		"airdb/scf-airdb",
		"airdb/scf-oauth2",
		"airdb/scf-wechat",
		"airdb/scf-mina",
		"airdb/scf-svgo",
	}

	c.String(http.StatusOK, strings.Join(modules, "\n"))
}

func (ctrl Controller) ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(bitbank, pair) // Dependency Injection
	c.JSON(200, ticker)
}

/*
func (ctrl Controller) candlestick(c *gin.Context) {
	args := usecase.OhlcArgs{
		E: bitbank, // Dependency Injection
		P: valueobject.BtcJpy,
		T: valueobject.OneMin,
	}
	candlestick := usecase.Ohlc(args)
	c.JSON(200, candlestick)
}

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(parameterRepository) // Dependency Injection
	c.JSON(200, parameter)
}

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewCardAndEatCheese(orderRepository) // Dependency Injection
	c.JSON(200, order)
}
*/

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewCardAndEatCheese(orderRepository) // Dependency Injection
	c.JSON(200, order)
}
