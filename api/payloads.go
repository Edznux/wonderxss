package api

import (
	"fmt"

	"github.com/edznux/wonderxss/crypto"
	"github.com/edznux/wonderxss/events"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func GetPayloads() ([]Payload, error) {
	fmt.Println("api.GetPayloads")
	fmt.Println(store)
	data, err := store.GetPayloads()
	if err != nil {
		return nil, err
	}
	payloads := []Payload{}
	for _, p := range data {
		tmp := Payload{}
		payloads = append(payloads, tmp.fromStorage(p))
	}

	return payloads, nil
}

func ServePayload(idOrAlias string) (string, error) {
	var err error
	var payload models.Payload

	payload, err = store.GetPayloadByAlias(idOrAlias)
	if err == models.NoSuchItem {
		// error will fallback (get overrided by the next call)
		payload, err = store.GetPayload(idOrAlias)
	}
	if err != nil {
		return "", err
	}
	// Run alert and store in DB without blocking.
	go func() {
		fmt.Println("=================================")
		fmt.Println("Notification should be sent now !")
		events.Events.Pub(payload, events.TOPIC_PAYLOAD_DELIVERED)
		fmt.Println("=================================")
		fmt.Println("Saving execution")
		AddExecution(payload.ID, idOrAlias)
		fmt.Println("=================================")
	}()
	return payload.Content, nil
}

func GetPayload(id string) (Payload, error) {
	payload, err := store.GetPayload(id)
	if err != nil {
		return Payload{}, err
	}
	res := Payload{}
	return res.fromStorage(payload), nil
}

//AddPayload is the API to add a new payload
func AddPayload(name string, content string) (models.Payload, error) {
	fmt.Printf("AddPayload(\"%s\", \"%s\")\n", name, content)
	hashes := crypto.GenerateSRIHashes(content)
	p := models.Payload{
		ID:      uuid.New().String(),
		Name:    name,
		Hash:    hashes.String(),
		Content: content,
	}
	fmt.Println(p)
	returnedPayload, err := store.CreatePayload(p)
	if err != nil {
		return models.Payload{}, err
	}

	return returnedPayload, nil
}
