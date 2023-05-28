/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "slackbot",
	Short: "Send content to Slack via Slack Application",
	Long: `SlackBot is an a application, which interacts with Slack application.
The application is able to send messages and attachments.
The recieved messages are fomratted, but the formatting settings are not implemented.`,
	Version: "v.0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	token, err := rootCmd.Flags().GetString("token")
	if err != nil {
		log.Fatal(err)
	}
	if token == "" {
		log.Fatal(`required flag(s) "token" not set`)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("token", "k", os.Getenv("SLACK_AUTH_TOKEN"), "Slack App Token")
	// rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.PersistentFlags().StringP("channel", "c", "", "Slack Workspace Channel Name")
	rootCmd.MarkPersistentFlagRequired("channel")

	versionTemplate := `{{printf "%s\n" .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}
