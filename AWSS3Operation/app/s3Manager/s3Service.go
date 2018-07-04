package s3Manager

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)

/*
S3Client ...
*/
type S3Client struct {
  SVCObject *s3.S3
  Location  string
}

/*
NewSVC ...
Location : https://docs.aws.amazon.com/zh_cn/general/latest/gr/rande.html
*/
func NewSVC(keyID, keySecret string, Location string) (client S3Client) {
  sess, _ := session.NewSession(&aws.Config{
    Region:      aws.String(Location),
    Credentials: credentials.NewStaticCredentials(keyID, keySecret, ""),
  })

  // Create S3 service client
  SVCObject := s3.New(sess)
  client = S3Client{SVCObject, Location}
  return
}

func haveErr(err error) (state bool) {
  state = true
  if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
      switch aerr.Code() {
      case s3.ErrCodeNoSuchBucket:
        fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
      default:
        fmt.Println(aerr.Error())
      }
    } else {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
    }
    return
  }
  state = false
  return
}
