package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const (
	bytesPerMegabyte = 1024 * 1024
	bytesPerGigabyte = 1024 * 1024 * 1024
)

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## GCP Storage")

	/*
		This code creates a client for interacting with Google Cloud Storage, then uses the client to list the objects in a bucket named "golang-snippets-bucket". It prints the name of each object to the console.

		If you test the code, be sure to change the bucket name accordingly and don't forget to provide the key file. For example, you can set the GOOGLE_APPLICATION_CREDENTIALS environment variable to the path of the JSON key file:

		export GOOGLE_APPLICATION_CREDENTIALS="/path/to/keyfile.json"
	*/

	ctx := context.Background()

	// Create a client for interacting with Google Cloud Storage.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List the objects in the "golang-snippets-bucket" bucket.
	iter := client.Bucket("golang-snippets-bucket").Objects(ctx, nil)
	for {
		obj, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to list objects: %v", err)
		}

		/*
			This code divides the size of each object by the appropriate number of bytes per megabyte or gigabyte, depending on the size of the object. It then prints the size in the appropriate unit (megabytes or gigabytes).
		*/
		size := obj.Size
		sizeUnit := "bytes"
		if size > bytesPerGigabyte {
			size = size / bytesPerGigabyte
			sizeUnit = "GB"
		} else if size > bytesPerMegabyte {
			size = size / bytesPerMegabyte
			sizeUnit = "MB"
		}
		fmt.Printf("%s (%d %s)\n", obj.Name, size, sizeUnit)
	}
}
