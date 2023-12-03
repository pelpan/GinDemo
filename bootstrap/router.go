package bootstrap

import (
	"fmt"
	"strconv"

	"GinDemo/routes"
	"GinDemo/utils"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/apiDemo")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

func RunServer() {
	port := 8080
	availPort, err := utils.FindAvailablePort(port)

	if err != nil {
		fmt.Println("Failed to find available port")
	}

	router := setupRouter()

	router.Run(":" + strconv.Itoa(availPort))
}
