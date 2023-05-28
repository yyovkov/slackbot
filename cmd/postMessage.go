/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"slackbot/slack"

	"github.com/spf13/cobra"
)

// postMessageCmd represents the sendMessage command
var postMessageCmd = &cobra.Command{
	Use:          "message",
	Short:        "post message to slack",
	Long:         `post message to Slack by specifying destination, title and text`,
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

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}

		text, err := cmd.Flags().GetString("text")
		if err != nil {
			return err
		}

		return sendMessage(title, text, token, channelID)
	},
}

func sendMessage(title, text, token, channelID string) error {
	msg := &slack.Message{}
	msg.Title = title
	msg.Text = text

	_, err := msg.Send(token, channelID)
	return err
}

func init() {
	postCmd.AddCommand(postMessageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	postMessageCmd.Flags().StringP("title", "t", "", "Message Title")
	postMessageCmd.MarkFlagRequired("title")

	postMessageCmd.Flags().StringP("text", "m", "", "Message text")
	postMessageCmd.MarkFlagRequired("text")
}
