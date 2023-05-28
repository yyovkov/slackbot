package slack

import (
	"github.com/slack-go/slack"
)

type File struct {
	Title          string
	InitialComment string
	Filetype       string
	Filename       string
	Channels       []string
	File           string
}

func (f *File) Upload(token, channelID string) (string, error) {
	// TODO: Check if file exists
	client := slack.New(token, slack.OptionDebug(false))

	fileUploadParams := slack.FileUploadParameters{
		Channels:       []string{channelID},
		File:           f.File,
		Filename:       f.Filename,
		Filetype:       f.Filetype,
		InitialComment: f.InitialComment,
		Title:          f.Title,
	}

	_, err := client.UploadFile(fileUploadParams)
	if err != nil {
		return "error", err
	}

	return "success", nil
}
