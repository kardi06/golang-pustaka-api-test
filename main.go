package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/books/:id", booksHandler)

	// query string
	// http://localhost:8888/query?name=bang&age=20
	router.GET("/query", queryHandler)

	// router.Run()
	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Bang Kar",
		"bio":  "Software Engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"subtitle": "Hello World",
		"title": "Belajar golang aja deh",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func queryHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age": age,
	})
}