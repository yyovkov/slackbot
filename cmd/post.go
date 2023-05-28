/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post Slack Channel",
	Long: `Post to Slack channels with predefined parameters
	* message - sending message
	* file - delete message`,
}

func init() {
	rootCmd.AddCommand(postCmd)
}
