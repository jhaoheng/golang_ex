package minio

import (
  // "fmt"
  miniogo "github.com/minio/minio-go"
)

/*
GetObject ...
*/
func GetObject(minioClient *miniogo.Client, bucket string, objectName string) (object *miniogo.Object, err error) {

  object, err = minioClient.GetObject(bucket, objectName, miniogo.GetObjectOptions{})
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  return object, err

  // localFile, err := os.Create("/tmp/local-file.jpg")
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  // if _, err = io.Copy(localFile, object); err != nil {
  //   fmt.Println(err)
  //   return
  // }
}
