package local

import (
	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/crypto"
	"github.com/edznux/wonderxss/events"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (local *Local) GetPayloads() ([]api.Payload, error) {
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
		log.Debugf("Notification should be sent now !")
		events.Events.Pub(payload, events.TOPIC_PAYLOAD_DELIVERED)
		log.Debugln("Saving execution")
		local.AddExecution(payload.ID, idOrAlias)
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
	log.Debugf("AddPayload(\"%s\", \"%s\", \"%s\")\n", name, content, contentType)
	hashes := crypto.GenerateSRIHashes(content)
	p := models.Payload{
		ID:          uuid.New().String(),
		Name:        name,
		Hashes:      hashes,
		Content:     content,
		ContentType: contentType,
	}
	log.Debugln(p)
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
