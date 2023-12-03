package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GinDemo/model"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/getTest", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": name,
		})
	})

	router.POST("/postTest", func(c *gin.Context) {
		name := c.PostForm("name")
		bankNoArr := c.PostFormArray("bankNo")

		var ret []string

		ret = append(ret, bankNoArr...)

		result := make(map[string]interface{})
		result["name"] = name
		result["bankNos"] = ret

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": result,
		})

	})

	router.POST("postJsonTest", func(c *gin.Context) {
		var requestBody *model.RequestsBody

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
}
