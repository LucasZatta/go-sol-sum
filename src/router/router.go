package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sol-sum/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/maxsum", func(c *gin.Context) {
		controllers.MaxSum(c)
	})

	return router
}
