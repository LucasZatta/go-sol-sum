package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sol-sum/models"
)

func MaxSum(c *gin.Context) {
	var list models.ListRequest

	//Receiving the request body
	bindErr := c.ShouldBindJSON(&list)

	//error handling
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "message": "Request body not valid"})
		return
	}

	//Calling algorithm
	maxSum := models.Kadanes(list.List)

	c.JSON(http.StatusOK, gin.H{"result": maxSum})
}
