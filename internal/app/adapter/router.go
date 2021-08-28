package adapter

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/airdb/sailor/deployutil"
	"github.com/airdb/sailor/version"
	"github.com/airdb/sls-default/internal/app/application/usecase"
	service "github.com/airdb/sls-default/internal/app/domain/service"
	"github.com/airdb/sls-default/internal/app/domain/valueobject"

	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"

	"github.com/gin-gonic/gin"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/airdb/sls-default/docs"
)

var (
	bitbank = service.Bitbank{}
	user    = service.Bitbank{}
	// parameterRepository = repository.Parameter{}
	// orderRepository     = repository.Order{}
)

// Controller is a controller
type Controller struct{}

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return GinFaas.ProxyWithContext(ctx, req)
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

	// Init the loc and set timezone.
	version.Init()

	r.GET("/", DefaultRoot)
	r.GET(projectPath, DefaultString)

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/hello", index)

	r.GET("/ticker", ticker)
	r.GET("/user/query", getUser)

	if deployutil.IsStageDev() {
		defaultAddr := ":8081"
		err := r.Run(defaultAddr)
		if err != nil {
			panic(err)
		}

		return
	}

	GinFaas = ginAdapter.New(r)

	faas.Start(Handler)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello airdb",
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

func ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(bitbank, pair) // Dependency Injection
	c.JSON(http.StatusOK, ticker)
}

func getUser(c *gin.Context) {
	user := usecase.GetUser(user)
	c.JSON(http.StatusOK, user)
}
