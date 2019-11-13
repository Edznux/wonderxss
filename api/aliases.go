package api

import (
	"fmt"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func GetAliases() ([]Alias, error) {
	data, err := store.GetAliases()
	if err != nil {
		return nil, err
	}
	aliases := []Alias{}
	for _, p := range data {
		tmp := Alias{}
		aliases = append(aliases, tmp.fromStorage(p))
	}

	return aliases, nil
}

func GetAlias(id string) (Alias, error) {
	alias, err := store.GetAlias(id)
	if err != nil {
		return Alias{}, err
	}
	res := Alias{}
	return res.fromStorage(alias), nil
}

func AddAlias(name string, payloadId string) (models.Alias, error) {
	fmt.Printf("AddAlias(\"%s\", \"%s\")\n", name, payloadId)
	p := models.Alias{
		ID:        uuid.New().String(),
		Short:     name,
		PayloadID: payloadId,
	}
	fmt.Println(p)
	returnedAlias, err := store.CreateAlias(p)
	if err != nil {
		return models.Alias{}, err
	}

	return returnedAlias, nil
}
