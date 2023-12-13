package bootstrap

import (
	"fmt"
	"strconv"

	"GinDemo/routes"
	"GinDemo/utils"

	"github.com/gin-gonic/gin"

	docs "GinDemo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/apiDemo")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /apiDemo
func RunServer() {
	port := 8080
	availPort, err := utils.FindAvailablePort(port)

	if err != nil {
		fmt.Println("Failed to find available port")
	}

	router := setupRouter()

	// swagger docs config
	docs.SwaggerInfo.BasePath = "/apiDemo"
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + strconv.Itoa(availPort))
}
