/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "slackbot/slack"
	"slackbot/slack"

	"github.com/spf13/cobra"
)

// postFileCmd represents the postFile command
var postFileCmd = &cobra.Command{
	Use:          "file",
	Short:        "Upload file to slack channel",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		token, err := cmd.Flags().GetString("token")
		if err != nil {
			return err
		}

		channelID, err := cmd.Flags().GetString("channel")
		if err != nil {
			return err
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}

		initialComment, err := cmd.Flags().GetString("initialComment")
		if err != nil {
			return err
		}

		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}

		f := &slack.File{}
		f.Filename = filename
		f.File = file
		f.InitialComment = initialComment
		f.Title = title

		return uploadFile(*f, token, channelID)
	},
}

func uploadFile(file slack.File, token, channelID string) error {
	// _, err := file.Upload(token, channelID)
	_, err := file.UploadFromS3(token, channelID)
	return err
}

func init() {
	postCmd.AddCommand(postFileCmd)

	postFileCmd.Flags().StringP("title", "t", "", "Title of the file to be displayed in the message")
	postFileCmd.Flags().StringP("initialComment", "i", "", "Initial comment to the uploaded file")
	postFileCmd.Flags().StringP("filetype", "e", "", "File type (Optional)")
	postFileCmd.Flags().StringP("filename", "n", "", "File type (Optional)")
	postFileCmd.Flags().StringP("file", "f", "", "Path to file for upload")

	postFileCmd.MarkFlagRequired("file")
}
