package oss

import (
  "app/config"
  // "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  // "os"
)

/*
DeleteObject ...
*/
func DeleteObject(objectName string) {

  data := config.AliyunOSS()
  Endpoint := data["aliyun_Endpoint"]
  AccessKeyID := data["aliyun_AccessKeyId"]
  AccessKeySecret := data["aliyun_AccessKeySecret"]
  Bucket := data["aliyun_Bucket"]

  client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
  if err != nil {
    // handleError(err)
  }
  bucket, err := client.Bucket(Bucket)
  if err != nil {
    // handleError(err)
  }
  err = bucket.DeleteObject(objectName)
  if err != nil {
    // handleError(err)
  }
}

// func handleError(err error) {
//   fmt.Println("Error:", err)
//   os.Exit(-1)
// }
