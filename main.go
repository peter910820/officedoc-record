package main

import (
	// "fmt"
	// "gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	// "officedoc-record/models"
)

func main() {
	r := gin.Default()

	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		officialDocument := c.PostForm("official-document")
		destination := c.PostForm("destination")

		c.JSON(http.StatusOK, gin.H{
			"name":             name,
			"officialDocument": officialDocument,
			"selectedOption":   destination,
		})

	})

	r.Run(":5000")
}
