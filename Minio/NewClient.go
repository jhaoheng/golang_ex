package minio

import (
  "app/config"
  "fmt"
  "github.com/minio/minio-go"
)

/*
NewClient ...
*/
func NewClient() *minio.Client {

  data := config.Minio()
  endpoint := data["minio_endpoint"]
  accessKeyID := data["minio_accessKeyID"]
  secretAccessKey := data["minio_secretAccessKey"]
  var useSSL bool
  if data["minio_useSSL"] == "true" {
    useSSL = true
  } else {
    useSSL = false
  }

  minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  if err != nil {
    fmt.Println(err)
  }
  return minioClient
}
