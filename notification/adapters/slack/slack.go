package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/edznux/wonder-xss/config"
	"github.com/edznux/wonder-xss/notification/interfaces"
)

type Slack struct {
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

func New(config config.Config) interfaces.NotificationSystem {
	s := Slack{}
	return &s
}

func (s *Slack) SendMessage(data string, destination string) error {
	webhookURL := "https://hooks.slack.com/services/" + destination
	slackBody, _ := json.Marshal(SlackRequestBody{Text: data})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
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
