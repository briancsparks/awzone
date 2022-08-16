package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"

  "github.com/spf13/cobra"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
  Use:   "s3",
  Short: "Manipulate S3 objects",

  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("s3 called")
  },
}

func init() {
  rootCmd.AddCommand(s3Cmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // s3Cmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // s3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
