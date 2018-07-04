package minio

import (
  "fmt"
  miniogo "github.com/minio/minio-go"
  "net/url"
  "strings"
  "time"
)

/*
PresignedGetObject ...
*/
func PresignedGetObject(minioClient *miniogo.Client, bucket string, objectName string) {

  objectComponents := strings.Split(objectName, "/")
  // fmt.Println(objectComponents)
  // fmt.Println(objectComponents[len(objectComponents)-1:])
  // file := objectComponents[len(objectComponents)-1:]
  file := objectComponents[len(objectComponents)-1]
  // fmt.Println(file)
  // return

  // Set request parameters for content-disposition.
  reqParams := make(url.Values)
  reqParams.Set("response-content-disposition", "attachment; filename=\""+file+"\"")

  // Generates a presigned url which expires in a day.
  presignedURL, err := minioClient.PresignedGetObject(bucket, objectName, time.Second*60, reqParams)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(presignedURL)
}
