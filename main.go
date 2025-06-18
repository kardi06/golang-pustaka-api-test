package main

import (
	"github.com/gin-gonic/gin"
)

func main {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Bang Kar",
			"bio":  "Software Engineer",
		})
	})
}