package main

import (
  "app/s3Manager"
  // "bufio"
  "fmt"
  _ "os"
)

func main() {

  Location := "us-west-1"
  BucketName := Location + "-onlyfortest"

  yourKeyID := ""
  yourKeySecret := ""
  fmt.Println("Info ", "\n Location : "+Location, "\n BucketName : "+BucketName)
  fmt.Println(" key id : "+yourKeyID, "\n key secret : "+yourKeySecret)

  client := s3Manager.NewSVC(yourKeyID, yourKeySecret, Location)
  fmt.Println("\n", client, "\n====================")

  /* Bucket List */
  client.ListBuckets()

  /* Bucket Create */
  // client.CreateBucket(BucketName)

  /* Bucket Delete */
  // client.DeleteBucket(BucketName)

  /* Object List */
  // client.ListObjects(BucketName)

  /* Object Put */
  // f, _ := os.Open("testobj.png")
  // Key := "test.png"
  // client.PutObject(BucketName, Key, f)

  /* Object Delete */
  // keys := []string{"test.png"}
  // client.DeleteOjects(BucketName, keys)

  /* Object GetObjectSignedURL*/
  // TTLSec := 60
  // key := "test.png" // need full name
  // client.GetObjectSignedURL(BucketName, key, TTLSec)
}
