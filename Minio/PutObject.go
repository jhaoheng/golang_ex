package minio

import (
	// "bytes"

	"io"

	miniogo "github.com/minio/minio-go"
	// "io/ioutil"
	// "reflect"
	"log"
)

/*
PutObject ...
*/
func PutObject(minioClient *miniogo.Client, bucket string, objectName string, newObj io.Reader, size int64) {

	// fmt.Println(getSize(newObj))
	// return
	// size := getSize(newObj)
	// fmt.Println(size)
	// fmt.Println(reflect.TypeOf(size))
	// return

	// b, err := ioutil.ReadAll(newObj)
	// if err != nil {
	//   log.Println(err)
	// }
	// fmt.Println(len(b))

	// buf := new(bytes.Buffer)
	// fmt.Println(buf)
	// buf.ReadFrom(newObj)
	// size := int64(buf.Len())
	// fmt.Println(size)
	// defer buf.Close()
	// buf.Reset()

	n, err := minioClient.PutObject(bucket, objectName, newObj, size, miniogo.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully uploaded bytes: ", n)
}

// func getSize(stream io.Reader) int64 {
//   buf := new(bytes.Buffer)
//   buf.ReadFrom(stream)
//   return int64(buf.Len())
// }
