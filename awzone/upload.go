package awzone

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

var client *s3.Client

func uploadFiles(filenames []string, bucket, path string) error {

	if client == nil {
		// Load the Shared AWS Configuration (~/.aws/config)
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// Create an Amazon S3 service client
		client = s3.NewFromConfig(cfg)
	}

	numFiles := len(filenames)
	for i, filename := range filenames {
		uploadFile(client, filename, bucket, path, i+1, numFiles)
	}

	return nil
}

func uploadFile(client *s3.Client, filename, bucket, path string, x, y int) error {
	// Open the file for use
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)

	count, err := file.Read(buffer)
	if err != nil {
		return err
	}
	_ = count

	basename := filepath.Base(filename)
	ext := filepath.Ext(filename)
	contentType := http.DetectContentType(buffer)
	if isBoringMimeType(contentType) {
		contentType = mime.TypeByExtension(ext)
	}

	var r io.Reader = bytes.NewReader(buffer)
	if size > 1024*1024 {
		r = io.TeeReader(r, NewProgressCounter(size, x, y, "Uploading "+basename+"..."))
	}

	output, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", path, basename)),
		//ACL:                "",
		Body:               r,
		ContentLength:      *aws.Int64(size),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		return err
	}
	_ = output

	//fmt.Printf("output: %v, ETag: %s\n", output, *output.ETag)

	return nil
}
