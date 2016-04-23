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

	router.GET("/", wfeIndex)
	router.POST("/auth", userAuth)

	router.Run(":8080")
}

func wfeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func userAuth(c *gin.Context) {
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
			if *resp.Item["Email"].S == form.Email && *resp.Item["Password"].S == form.Password {
				// c.JSON(200, gin.H{"status": "Logged in!"})
				bucketlist := ""
				for _, dataset := range resp.Item["Datasets"].SS {
					bucketlist = bucketlist + makeListItem(*dataset)
					// s3instance := s3.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
					// page := 0
					// err := s3instance.ListObjectsPages(&s3.ListObjectsInput{
					// 	Bucket: dataset,
					// }, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
					// 	fmt.Println("Page,", page)
					// 	for _, obj := range p.Contents {
					// 		fmt.Println("Object:", *obj.Key)
					// 	}
					// 	return true
					// })
					// if err != nil {
					// 	fmt.Println("failted to list objects", err)
					// }
				}
				fmt.Println(bucketlist)
				c.HTML(http.StatusOK, "bucketlist.tmpl", gin.H{
					"bucketlist": bucketlist,
				})
			} else {
				// c.JSON(401, gin.H{"status": "Unauthorized"})
				c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
					"message": "Invalid login information.",
				})
			}
		}
	}
}

func makeListItem(bucket string) string {
	liststring := "<li>" + bucket + "</li>"
	return liststring
}
