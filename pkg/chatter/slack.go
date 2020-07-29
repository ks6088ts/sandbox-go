package chatter

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Session ...
type Session struct {
	client *slack.Client
}

// NewSession creates a Session
func NewSession(token string) (*Session, error) {
	return &Session{
		client: slack.New(token),
	}, nil
}

// PostMessage ...
func (s *Session) PostMessage(message string, channelID string, iconURL string) error {
	channelID, timestamp, err := s.client.PostMessage(channelID, slack.MsgOptionText(message, false), slack.MsgOptionIconURL(iconURL))
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
