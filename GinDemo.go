package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net"
	"os"
	"strconv"
)

type RequestsBody struct {
	Name    string   `json:"name"`
	BankNos []string `json:"bankNos"`
}

func checkPortAvailability(port int) bool {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	defer listener.Close()
	return true
}

func findAvailablePort(port int) (int, error) {
	for startPort := port + 1; startPort < 65535; startPort++ {
		fmt.Println("checking port: ", startPort)
		if checkPortAvailability(startPort) {
			return startPort, nil
		}
	}

	return 0, fmt.Errorf("no available port")
}

func main() {
	port := 8080
	availPort, err := findAvailablePort(port)
	if err != nil {
		fmt.Println("Failed to find available port")
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/getTest", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": name,
		})
	})

	r.POST("/postTest", func(c *gin.Context) {
		name := c.PostForm("name")
		bankNoArr := c.PostFormArray("bankNo")

		var ret []string

		for _, bankNo := range bankNoArr {
			ret = append(ret, bankNo)
		}

		result := make(map[string]interface{})
		result["name"] = name
		result["bankNos"] = ret

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": result,
		})

	})

	r.POST("postJsonTest", func(c *gin.Context) {
		var requestBody RequestsBody

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := make(map[string]interface{})
		result["name"] = requestBody.Name
		result["bankNos"] = requestBody.BankNos

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": result,
		})

	})

	startErr := r.Run(":" + strconv.Itoa(availPort))
	if startErr != nil {
		fmt.Println("Failed to start server")
		return
	}
}
