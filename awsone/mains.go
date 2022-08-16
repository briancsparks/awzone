package awsone

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

func UploadMain(filenames []string, bucket, path string) error {

	return uploadFiles(filenames, bucket, path)
}
