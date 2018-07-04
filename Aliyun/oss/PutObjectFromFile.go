package oss

import (
  "app/config"
  "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  // "os"
)

/*
PutObjectFromFile ...
*/
func PutObjectFromFile() {
  data := config.AliyunOSS()
  Endpoint := data["aliyun_Endpoint"]
  AccessKeyID := data["aliyun_AccessKeyId"]
  AccessKeySecret := data["aliyun_AccessKeySecret"]
  Bucket := data["aliyun_Bucket"]

  client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
  if err != nil {
    // HandleError(err)
  }

  bucket, err := client.Bucket(Bucket)
  if err != nil {
    // HandleError(err)
  }

  err = bucket.PutObjectFromFile("dog test", "https://vod-shenzhen.oss-cn-shenzhen.aliyuncs.com/my-object?spm=5176.8466032.bucket-object.dopenurl.25821450mBGppN&Expires=1525760505&OSSAccessKeyId=TMP.AQEvDqHzuDbe67SAZcfUn-LIcEf0WUhl_xE_bK0DZ7tjhQXViJTmtcbANo-7ADAtAhUAmQ9dBEayY4R-PZQ3kCg_ufGP9KsCFFhUwJ2u57siCqGA5soSa3TU7mh5&Signature=cXAWdh56hDBAWrfXFylvdNN6WzY%3D")
  if err != nil {
    fmt.Println(err)
    // HandleError(err)
  }
}
