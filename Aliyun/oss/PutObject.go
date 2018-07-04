package oss

import (
  "app/config"
  // "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  "io"
  // "strings"
  "log"
)

/*
PutObject ...
*/
func PutObject(name string, newObj io.Reader) error {
  data := config.AliyunOSS()
  Endpoint := data["aliyun_Endpoint"]
  AccessKeyID := data["aliyun_AccessKeyId"]
  AccessKeySecret := data["aliyun_AccessKeySecret"]
  Bucket := data["aliyun_Bucket"]

  client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
  if err != nil {
    log.Println(err)
    return err
    // handleError(err)
  }
  bucket, err := client.Bucket(Bucket)
  if err != nil {
    log.Println(err)
    // handleError(err)
    return err
  }
  err = bucket.PutObject(name, newObj)
  if err != nil {
    log.Println(err)
    // handleError(err)
    return err
  }
  return nil
}

// func handleError(err error) {
//   fmt.Println("Error:", err)
//   os.Exit(-1)
// }
