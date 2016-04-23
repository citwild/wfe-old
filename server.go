package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

// Login is a struct to get form data from website
type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.POST("/auth", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			dbInstance := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
			params := &dynamodb.GetItemInput{
				Key: map[string]*dynamodb.AttributeValue{ // Required
					"Email": { // Required
						S: aws.String(form.Email),
					},
				},
				TableName:      aws.String("WFE"),
				ConsistentRead: aws.Bool(true),
			}
			resp, err := dbInstance.GetItem(params)
			if err != nil {
				c.JSON(401, gin.H{"status": "DB error"})
			} else {
				fmt.Println(*resp.Item["Password"].S)
				// if *resp.Item["Email"].S == form.Email && *resp.Item["Password"].S == form.Password {
				// 	c.JSON(200, gin.H{"status": "Logged in!"})
				// } else {
				// 	c.JSON(401, gin.H{"status": "Unauthorized"})
				// }
			}
		}
	})

	router.Run(":8080")
}
