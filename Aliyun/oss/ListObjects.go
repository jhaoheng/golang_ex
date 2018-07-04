package oss

import (
  "app/config"
  // "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  // "os"
)

/*
ListObjects ...
*/
func ListObjects() []oss.ObjectProperties {
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
  lsRes, err := bucket.ListObjects()
  if err != nil {
    // handleError(err)
  }
  // for _, object := range lsRes.Objects {
  //   fmt.Println("Object:", object.Key)
  // }
  return lsRes.Objects
}

// func handleError(err error) {
//   fmt.Println("Error:", err)
//   os.Exit(-1)
// }
