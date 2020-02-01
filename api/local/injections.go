package local

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func (local *Local) GetInjections() ([]api.Injection, error) {
	data, err := local.store.GetInjections()
	if err != nil {
		return nil, err
	}
	injections := []api.Injection{}
	for _, p := range data {
		tmp := api.Injection{}
		injections = append(injections, tmp.FromStorage(p))
	}

	return injections, nil
}

func (local *Local) GetInjection(id string) (api.Injection, error) {
	injection, err := local.store.GetInjection(id)
	if err != nil {
		return api.Injection{}, err
	}
	res := api.Injection{}
	return res.FromStorage(injection), nil
}

func (local *Local) AddInjection(name string, content string) (api.Injection, error) {
	var returnedInjection api.Injection
	fmt.Printf("AddInjection(\"%s\", \"%s\")\n", name, content)
	p := models.Injection{
		ID:      uuid.New().String(),
		Name:    name,
		Content: content,
	}
	injection, err := local.store.CreateInjection(p)
	if err != nil {
		log.Println("could not add Injection:", err)
		return api.Injection{}, err
	}

	return returnedInjection.FromStorage(injection), nil
}

func (local *Local) DeleteInjection(id string) error {
	e := models.Injection{ID: id}
	err := local.store.DeleteInjection(e)
	if err != nil {
		return err
	}
	return nil
}
