package local

import (
	"fmt"
	"log"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func (local *Local) GetAliases() ([]api.Alias, error) {
	data, err := local.store.GetAliases()
	if err != nil {
		return nil, err
	}
	aliases := []api.Alias{}
	for _, p := range data {
		tmp := api.Alias{}
		aliases = append(aliases, tmp.FromStorage(p))
	}

	return aliases, nil
}

func (local *Local) GetAlias(id string) (api.Alias, error) {
	alias, err := local.store.GetAlias(id)
	if err != nil {
		return api.Alias{}, err
	}
	res := api.Alias{}
	return res.FromStorage(alias), nil
}

func (local *Local) GetAliasByID(id string) (api.Alias, error) {
	alias, err := local.store.GetAliasByID(id)
	if err != nil {
		return api.Alias{}, err
	}
	res := api.Alias{}
	return res.FromStorage(alias), nil
}

func (local *Local) GetAliasByPayloadID(id string) (api.Alias, error) {
	alias, err := local.store.GetAliasByPayloadID(id)
	if err != nil {
		return api.Alias{}, err
	}
	res := api.Alias{}
	return res.FromStorage(alias), nil
}

func (local *Local) AddAlias(name string, payloadId string) (api.Alias, error) {
	fmt.Printf("AddAlias(\"%s\", \"%s\")\n", name, payloadId)
	var returnedAlias api.Alias
	p := models.Alias{
		ID:        uuid.New().String(),
		Short:     name,
		PayloadID: payloadId,
	}
	alias, err := local.store.CreateAlias(p)
	if err != nil {
		log.Println("could not add alias:", err)
		return api.Alias{}, err
	}

	return returnedAlias.FromStorage(alias), nil
}

func (local *Local) DeleteAlias(id string) error {
	e := models.Alias{ID: id}
	err := local.store.DeleteAlias(e)
	if err != nil {
		return err
	}
	return nil
}
