package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()
	// This client will be used to connect to the Google Cloud Storage service.
	// Defaults variables comming from gcloud auth
	// If not works it is possible use package option to import credentials
	// option.WithCredentialsFile("path/to/credentials.json")
	client, err := storage.NewClient(ctx)

	fileName := "test.txt"
	bucketName := "test-bucket"
	rc, err := client.Bucket(bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()

	body, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
