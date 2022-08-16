package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/briancsparks/awsone/awsone"

	"github.com/spf13/cobra"
)

var bucket, path string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to S3",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called", args)
		awsone.UploadMain([]string{args[0]}, bucket, path)
	},
}

func init() {
	s3Cmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	uploadCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "The bucket to upload to")
	uploadCmd.Flags().StringVarP(&path, "path", "p", "", "The path part of the S3 key")
}
