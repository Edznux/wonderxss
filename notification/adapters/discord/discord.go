package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/edznux/wonderxss/events"
	"github.com/edznux/wonderxss/storage/models"
)

type Discord struct {
	Name string
}

type DiscordRequestBody struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func New(config Config) *Discord {
	fmt.Println("New Notification Handler: discord")
	s := Discord{Name: "WonderXSS"}
	ch := events.Events.Sub(events.TOPIC_PAYLOAD_DELIVERED)

	go func(ch chan interface{}) {
		for {
			if msg, ok := <-ch; ok {
				fmt.Printf("Received %s, times.\n", msg)
				payload := msg.(models.Payload)
				notif := "A payload was triggered : " + payload.Name + " at " + time.Now().String()
				s.sendMessage(notif, config.WebHookURL)
			} else {
				fmt.Println("not ok")
				break
			}
		}
	}(ch)

	return &s
}

func (s *Discord) sendMessage(data string, destination string) error {
	discordBody, _ := json.Marshal(DiscordRequestBody{
		Content:  data,
		Username: s.Name,
	})
	req, err := http.NewRequest(http.MethodPost, destination, bytes.NewBuffer(discordBody))
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
		return errors.New("Non-ok response returned from Discord")
	}
	return nil
}
