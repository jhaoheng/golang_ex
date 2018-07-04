package s3Manager

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "io"
  "time"
)

/*
ListObjects ...
*/
func (client S3Client) ListObjects(BucketName string) {
  input := &s3.ListObjectsInput{
    Bucket:  aws.String(BucketName),
    MaxKeys: aws.Int64(2),
  }
  result, err := client.SVCObject.ListObjects(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result.Contents)
}

/*
PutObject ...
*/
func (client S3Client) PutObject(BucketName, Key string, newObj io.Reader) {
  input := &s3.PutObjectInput{
    Body:   aws.ReadSeekCloser(newObj),
    Bucket: aws.String(BucketName),
    Key:    aws.String(Key),
  }
  result, err := client.SVCObject.PutObject(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result)
}

/*
GetObject ...
*/
func (client S3Client) GetObject(BucketName string, key string) {
  input := &s3.GetObjectInput{
    Bucket: aws.String(BucketName),
    Key:    aws.String(key),
  }
  result, err := client.SVCObject.GetObject(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result)
}

/*
GetObjectSignedURL ...
*/
func (client S3Client) GetObjectSignedURL(BucketName string, key string, TTLSec int) {
  req, _ := client.SVCObject.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(BucketName),
    Key:    aws.String(key),
  })
  urlStr, _ := req.Presign(time.Duration(TTLSec) * time.Second)
  fmt.Println(urlStr)
}

/*
DeleteOjects ...
*/
func (client S3Client) DeleteOjects(BucketName string, keys []string) {

  var objs = make([]*s3.ObjectIdentifier, len(keys))
  for i, v := range keys {
    // Add objects from command line to array
    objs[i] = &s3.ObjectIdentifier{Key: aws.String(v)}
  }

  input := &s3.DeleteObjectsInput{
    Bucket: aws.String(BucketName),
    Delete: &s3.Delete{
      Objects: objs,
      Quiet:   aws.Bool(false),
    },
  }
  result, err := client.SVCObject.DeleteObjects(input)
  if haveErr(err) {
    return
  }
  fmt.Println(result)
}
