package local

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// GetLoots return all the triggered payload stored in the database.
func (local *Local) GetLoots() ([]api.Loot, error) {
	data, err := local.store.GetLoots()
	if err != nil {
		return nil, err
	}
	loots := []api.Loot{}
	log.Debugln("Loots from store: ", data)
	for _, p := range data {
		tmp := api.Loot{}
		loots = append(loots, tmp.FromStorage(p))
	}

	return loots, nil
}

func (local *Local) GetLoot(id string) (api.Loot, error) {
	loot, err := local.store.GetLoot(id)
	if err != nil {
		return api.Loot{}, err
	}
	res := api.Loot{}
	return res.FromStorage(loot), nil
}

func (local *Local) AddLoot(data string) (api.Loot, error) {
	var returnedLoot api.Loot
	fmt.Printf("AddLoot(\"%s\")\n", data)
	c := models.Loot{
		ID:        uuid.New().String(),
		Data:      data,
		CreatedAt: time.Now(),
	}
	fmt.Println(c)
	loot, err := local.store.CreateLoot(c)
	if err != nil {
		return api.Loot{}, err
	}

	return returnedLoot.FromStorage(loot), nil
}

func (local *Local) DeleteLoot(id string) error {
	e := models.Loot{ID: id}
	err := local.store.DeleteLoot(e)
	if err != nil {
		return err
	}
	return nil
}
