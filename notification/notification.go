package notification

import (
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification/adapters/slack"
)

// var RegisteredNotificationSystem []interfaces.NotificationSystem

func Setup(cfg config.Config) {
	// RegisteredNotificationSystem = []interfaces.NotificationSystem{}
	for _, ns := range cfg.Notifications {
		if ns.Name == "slack" {
			_ = slack.New(slack.Config{WebHookURL: ns.Token})
		}
	}
	// RegisteredNotificationSystem = append(RegisteredNotificationSystem, slackNS)
}
