package main

import (
	"fmt"
	"net/http"

	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// Login is a struct to get login form data
type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// SelectedBucket is a struct to get the selected bucket form data
type SelectedBucket struct {
	Bucket string `form:"selectbucket" json:"selectbucket" binding:"required"`
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// TODO: requestccess route
func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	// Uncomment the next line when ready for release.
	// gin.SetMode(gin.ReleaseMode)

	router.GET("/", wfeIndex)
	router.GET("/contact", wfeContact)
	router.POST("/auth", userAuth)
	router.POST("/bucketlist", bucketShow)
	// router.POST("/requestccess", wfeRequestAccess)

	router.Run(":8080")
}

func wfeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func wfeContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.tmpl", gin.H{})
}

func userAuth(c *gin.Context) {
	var form Login
	Info.Println("Authorizing user")
	if c.Bind(&form) == nil {
		dbInstance := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
		params := &dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Email": {
					S: aws.String(form.Email),
				},
			},
			TableName:      aws.String("WFE"),
			ConsistentRead: aws.Bool(false),
		}
		resp, err := dbInstance.GetItem(params)
		if err != nil {
			Error.Println("Error getting item")
			c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
				"message": "Database error.",
			})
		} else {
			if _, ok := resp.Item["Email"]; ok {
				if *resp.Item["Email"].S == form.Email && *resp.Item["Password"].S == form.Password {
					var bucketlist []string
					for _, dataset := range resp.Item["Datasets"].SS {
						bucketlist = append(bucketlist, *dataset)
					}
					Info.Println("User password and email match")
					c.HTML(http.StatusOK, "bucketlist.tmpl", gin.H{
						"bucketlist": bucketlist,
					})
				} else {
					Info.Println("Failure authorizing user: Invalid login")
					c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
						"message": "Invalid login information.",
					})
				}
			} else {
				Info.Println("Failure authorizing user: Invalid login")
				c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
					"message": "Invalid login information.",
				})
			}
		}
	} else {
		Info.Println("Failure authorizing user: No input provided")
		c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
			"message": "Please fill the form with valid login information.",
		})
	}
}

// TODO: Make a webpage for bucket-listing
// TODO: Print output of bucket-listing on the webpage.
// TODO: Add links to open the files.
// TODO: Permissions of the files.
// TODO: Show error message on page without reloading the page.
func bucketShow(c *gin.Context) {
	var bucket SelectedBucket
	if c.Bind(&bucket) == nil {
		s3instance := s3.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
		page := 0
		err := s3instance.ListObjectsPages(&s3.ListObjectsInput{
			Bucket: aws.String(bucket.Bucket),
		}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
			fmt.Println("Page,", page)
			for _, obj := range p.Contents {
				fmt.Println("Object:", *obj.Key)
			}
			return true
		})
		if err != nil {
			Error.Println("Failed to list objects", err)
		}
	} else {
		Info.Println("No buckets selected")
	}
}
