package main

import (
	"net/http"

	"gin/bar"
	"gin/foo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong2",
		})
	})
	r.GET("/ping100", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": foo.Foo(),
		})
	})
	r.GET("/ping101", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": bar.Bar(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
