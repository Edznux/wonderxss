package notification

import (
	"os"

	"github.com/edznux/wonderxss/notification/adapters/slack"
)

// var RegisteredNotificationSystem []interfaces.NotificationSystem

func Setup() {
	// RegisteredNotificationSystem = []interfaces.NotificationSystem{}
	token := os.Getenv("SLACK_WEBHOOK")
	_ = slack.New(slack.Config{WebHookURL: token})
	// RegisteredNotificationSystem = append(RegisteredNotificationSystem, slackNS)
}
