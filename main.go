package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(bucket, filename string) error {

	// Open session
	sess := session.Must(session.NewSession())

	// Open file to upload
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Unable to open file %q, %v", filename, err)
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	return nil

}

func Download(bucket, filename string) error {

	// Open session
	sess := session.Must(session.NewSession())

	// Create new file to fullfill with bucket content
	// Will fail if file exists
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Unable to open file %q, %v", filename, err)
	}

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		return fmt.Errorf("Unable to download item %q, %v", filename, err)
	}
	return nil
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s [upload|download] bucket_name local_filename",
			os.Args[0])
		os.Exit(1)
	}

	action := os.Args[1]
	bucket := os.Args[2]
	filename := os.Args[3]

	switch action {
	case "upload":
		err := Upload(bucket, filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	case "download":
		err := Download(bucket, filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("Successfully downloaded %q from %q\n", filename, bucket)

	}

}
