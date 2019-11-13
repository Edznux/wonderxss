package notification

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification/adapters/slack"
	"github.com/edznux/wonderxss/notification/interfaces"
	"github.com/edznux/wonderxss/storage/models"
)

var RegisteredNotificationSystem []interfaces.NotificationSystem

func init() {
	RegisteredNotificationSystem = []interfaces.NotificationSystem{}
	slackNS := slack.New(config.Config{})
	RegisteredNotificationSystem = append(RegisteredNotificationSystem, slackNS)
}

// SendNotifications just send all alert to all notification systems.
func SendNotifications(payload models.Payload) {
	fmt.Println("Sending notifications", RegisteredNotificationSystem)
	msg := "A payload was triggered : " + payload.Name + " at " + time.Now().String()
	// TODO: from config
	dest := "SLACK_TOKEN_HERE"
	for _, ns := range RegisteredNotificationSystem {
		fmt.Printf("Sending notification : %s \nto: %s\n", msg, dest)
		ns.SendMessage(msg, dest)
	}
}
