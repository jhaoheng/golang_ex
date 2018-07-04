package oss

import (
  "app/config"
  "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  // "os"
)

/*
DeleteObjects ...
*/
func DeleteObjects(objectsName []string) {

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

  delRes, err := bucket.DeleteObjects(objectsName)
  if err != nil {
    // HandleError(err)
  }
  fmt.Println("Deleted Objects:", delRes.DeletedObjects)

  // // 不返回删除的结果
  // _, err = bucket.DeleteObjects([]string{"my-object-3", "my-object-4"},
  //   oss.DeleteObjectsQuiet(true))
  // if err != nil {
  //   // HandleError(err)
  // }
}

// func handleError(err error) {
//   fmt.Println("Error:", err)
//   os.Exit(-1)
// }
