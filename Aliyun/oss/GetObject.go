package oss

import (
  "app/config"
  "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  "io/ioutil"
  // "os"
)

/*
GetObject ...
*/
func GetObject(objectName string) []byte {

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

  body, err := bucket.GetObject("my-object")
  if err != nil {
    // handleError(err)
  }

  dataObject, err := ioutil.ReadAll(body)
  if err != nil {
    // handleError(err)
  }
  body.Close()

  fmt.Println("data:", string(dataObject))
  return dataObject
}

// func handleError(err error) {
//   fmt.Println("Error:", err)
//   os.Exit(-1)
// }
