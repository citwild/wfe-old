package bucket

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


