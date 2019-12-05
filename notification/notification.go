package notification

import (
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification/adapters/discord"
	"github.com/edznux/wonderxss/notification/adapters/slack"
)

func Setup() {
	cfg := config.Current
	for _, nsCfg := range cfg.Notifications {
		if nsCfg.Enabled {
			if nsCfg.Name == "slack" {
				_ = slack.New(slack.Config{WebHookURL: nsCfg.Token})
			}
			if nsCfg.Name == "discord" {
				_ = discord.New(discord.Config{WebHookURL: nsCfg.Token})
			}
		}
	}
}
