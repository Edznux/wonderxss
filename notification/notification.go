package notification

import (
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification/adapters/slack"
)

func Setup(cfg config.Config) {
	for _, nsCfg := range cfg.Notifications {
		if nsCfg.Name == "slack" {
			_ = slack.New(slack.Config{WebHookURL: nsCfg.Token})
		}
	}
}
