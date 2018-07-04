package oss

import (
  // "fmt"
  "app/config"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
GetObjectSignedURL ...
*/
func GetObjectSignedURL(ObjectName string) string {

  data := config.AliyunOSS()
  Endpoint := data["aliyun_Endpoint"]
  AccessKeyID := data["aliyun_AccessKeyId"]
  AccessKeySecret := data["aliyun_AccessKeySecret"]
  Bucket := data["aliyun_Bucket"]

  client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
  if err != nil {
    return ""
    // handleError(err)
  }
  bucket, err := client.Bucket(Bucket)
  if err != nil {
    return ""
    // handleError(err)
  }

  // 取得 signURL
  signedURL, err := bucket.SignURL(ObjectName, oss.HTTPGet, 60)
  if err != nil {
    return ""
    // HandleError(err)
  }
  return signedURL
}
