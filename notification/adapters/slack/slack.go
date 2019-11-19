package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/edznux/wonderxss/events"
	"github.com/edznux/wonderxss/notification/interfaces"
	"github.com/edznux/wonderxss/storage/models"
)

type Slack struct {
	Name string
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

func New(config Config) interfaces.NotificationSystem {
	fmt.Println("New Notification Handler: slack")
	s := Slack{Name: "Slack"}
	ch := events.Events.Sub(events.TOPIC_PAYLOAD_DELIVERED)

	go func(ch chan interface{}) {
		for {
			if msg, ok := <-ch; ok {
				fmt.Printf("Received %s, times.\n", msg)
				payload := msg.(models.Payload)
				notif := "A payload was triggered : " + payload.Name + " at " + time.Now().String()
				s.SendMessage(notif, config.WebHookURL)
			} else {
				fmt.Println("not ok")
				break
			}
		}
	}(ch)

	return &s
}

func (s *Slack) SendMessage(data string, destination string) error {
	slackBody, _ := json.Marshal(SlackRequestBody{Text: data})
	req, err := http.NewRequest(http.MethodPost, destination, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
