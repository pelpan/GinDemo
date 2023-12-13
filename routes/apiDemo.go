package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GinDemo/model"
)

// @BasePath /apiDemo
// apiDemo godoc
// @Summary apiDemo
// @Schemes
// @Description apiDemo
// @Tags apiDemo
// @Produce json
// @Success 200 {string} json "{"code": 0,"data": name}"
// @Router /getTest [get]
func getTest(c *gin.Context) {

	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": name,
	})
}

// @BasePath /apiDemo
// apiDemo godoc
// @Summary apiDemo
// @Schemes
// @Description apiDemo
// @Param name formData string false "name"
// @Param bankNo formData string false "bankNo"
// @Tags apiDemo
// @Produce json
// @Success 200 {object} model.Result{name=string, bankNos=[]string} "result"
// @Router /postTest [post]
func postTest(c *gin.Context) {
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
}

// @BasePath /apiDemo
// apiDemo godoc
// @Summary apiDemo
// @Schemes
// @Description apiDemo
// @Param data body string false "Data"
// @Tags apiDemo
// @Produce json
// @Success 200 {object} model.Result{name=string, bankNos=[]string} "result"
// @Router /postJsonTest [post]
func postJsonTest(c *gin.Context) {
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
}

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/getTest", getTest)
	router.POST("/postTest", postTest)
	router.POST("postJsonTest", postJsonTest)
}
