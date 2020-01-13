package local

import (
	"fmt"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/crypto"
	"github.com/edznux/wonderxss/events"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func (local *Local) GetPayloads() ([]api.Payload, error) {
	fmt.Println("api.GetPayloads")
	fmt.Println(local.store)
	data, err := local.store.GetPayloads()
	if err != nil {
		return nil, err
	}
	payloads := []api.Payload{}
	for _, p := range data {
		tmp := api.Payload{}
		payloads = append(payloads, tmp.FromStorage(p))
	}

	return payloads, nil
}

func (local *Local) ServePayload(idOrAlias string) (string, error) {
	var err error
	var payload models.Payload

	payload, err = local.store.GetPayloadByAlias(idOrAlias)
	if err == models.NoSuchItem {
		// error will fallback (get overrided by the next call)
		payload, err = local.store.GetPayload(idOrAlias)
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
		local.AddExecution(payload.ID, idOrAlias)
		fmt.Println("=================================")
	}()
	return payload.Content, nil
}

func (local *Local) GetPayload(id string) (api.Payload, error) {
	payload, err := local.store.GetPayload(id)
	if err != nil {
		return api.Payload{}, err
	}
	res := api.Payload{}
	return res.FromStorage(payload), nil
}

//AddPayload is the API to add a new payload
func (local *Local) AddPayload(name string, content string, contentType string) (api.Payload, error) {
	var returnedPayload api.Payload
	fmt.Printf("AddPayload(\"%s\", \"%s\", \"%s\")\n", name, content, contentType)
	hashes := crypto.GenerateSRIHashes(content)
	p := models.Payload{
		ID:          uuid.New().String(),
		Name:        name,
		Hashes:      hashes,
		Content:     content,
		ContentType: contentType,
	}
	fmt.Println(p)
	payload, err := local.store.CreatePayload(p)
	if err != nil {
		return api.Payload{}, err
	}

	return returnedPayload.FromStorage(payload), nil
}

func (local *Local) DeletePayload(id string) error {
	e := models.Payload{ID: id}
	err := local.store.DeletePayload(e)
	if err != nil {
		return err
	}
	return nil
}
