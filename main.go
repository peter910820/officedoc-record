package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	r := gin.Default()

	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":5000")
}
