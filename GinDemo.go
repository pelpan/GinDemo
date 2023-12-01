package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestsBody struct {
	Name    string   `json:"name"`
	BankNos []string `json:"bankNos"`
}

func main() {
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

	err := r.Run()
	if err != nil {
		return
	}
}
