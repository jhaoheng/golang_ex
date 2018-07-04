package minio

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

/*
MakeBucket ...
*/
func MakeBucket(minioClient *minio.Client, bucket, location string) {
	err := minioClient.MakeBucket(bucket, location)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully created mybucket.")
}

/*
ListBuckets ...
*/
func ListBuckets(minioClient *minio.Client) {
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

/*
RemoveBucket ...
*/
func RemoveBucket() {

}

/*
BucketExists ...
*/
func BucketExists(minioClient *minio.Client, mybucket string) bool {
	found, err := minioClient.BucketExists(mybucket)
	if err != nil {
		log.Println(err)
		return false
	}
	return found
}
