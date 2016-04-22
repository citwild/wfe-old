package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Wide-Field Ethnography",
		})
	})

	router.POST("/auth", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			log.Printf("Email: %s\t Password: %s\n", form.Email, form.Password)
		}
	})

	router.Run(":8080")
}
