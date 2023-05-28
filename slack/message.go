package slack

import (
	"github.com/slack-go/slack"
)

type Message struct {
	Title string
	Text  string
	Color string
}

func (sm *Message) Send(token, channelID string) (string, error) {
	client := slack.New(token, slack.OptionDebug(false))

	attachment := slack.Attachment{
		Title:    sm.Title,
		Text:     sm.Text,
		AuthorID: "test",
		Color:    "#36a64f",
	}

	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)

	return timestamp, err
}
