package s3Manager

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
)

/*
CreateBucket ...
BucketName : Should Location-BucketName. EX: ap-northeast-1-fortest
*/
func (client S3Client) CreateBucket(BucketName string) {
  input := &s3.CreateBucketInput{
    Bucket: aws.String(BucketName),
    CreateBucketConfiguration: &s3.CreateBucketConfiguration{
      LocationConstraint: aws.String(client.Location),
    },
  }
  result, err := client.SVCObject.CreateBucket(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result)
}

/*
DeleteBucket ...
*/
func (client S3Client) DeleteBucket(BucketName string) {
  input := &s3.DeleteBucketInput{
    Bucket: aws.String(BucketName),
  }
  result, err := client.SVCObject.DeleteBucket(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result)
}

/*
ListBuckets ...
*/
func (client S3Client) ListBuckets() {
  input := &s3.ListBucketsInput{}
  result, err := client.SVCObject.ListBuckets(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result.Buckets)
}
