package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NumbersRequest struct {
	Numbers []int `json:"numbers"`
}

func sumNumbers(c *gin.Context) {
	var req NumbersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var sum int
	for _, num := range req.Numbers {
		sum += num
	}

	c.JSON(http.StatusOK, gin.H{"result": sum})
}

func main() {
	r := gin.Default()
	r.POST("/", sumNumbers)
	fmt.Println("Server is listening on port 8080...")
	r.Run(":8080")
}
