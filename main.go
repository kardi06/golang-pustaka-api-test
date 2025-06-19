package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Bang Kar",
			"bio":  "Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"subtitle": "Hello World",
			"title": "Belajar golang aja deh",
		})
	})

	// router.Run()
	router.Run(":8888")
}