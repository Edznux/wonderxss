package api

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

// GetLoots return all the triggered payload stored in the database.
func GetLoots() ([]Loot, error) {
	fmt.Println("api.GetLoots")
	data, err := store.GetLoots()
	if err != nil {
		return nil, err
	}
	loots := []Loot{}
	fmt.Println("Loots from store: ", data)
	for _, p := range data {
		tmp := Loot{}
		loots = append(loots, tmp.fromStorage(p))
	}

	return loots, nil
}

func GetLoot(id string) (Loot, error) {
	loot, err := store.GetLoot(id)
	if err != nil {
		return Loot{}, err
	}
	res := Loot{}
	return res.fromStorage(loot), nil
}

func AddLoot(payloadID string, aliasID string) (models.Loot, error) {
	fmt.Printf("AddLoot(\"%s\", \"%s\")\n", payloadID, aliasID)
	l := models.Loot{
		ID:          uuid.New().String(),
		PayloadID:   payloadID,
		AliasID:     aliasID,
		TriggeredAt: time.Now(),
	}
	fmt.Println(l)
	returnedAlias, err := store.CreateLoot(l, aliasID)
	if err != nil {
		return models.Loot{}, err
	}

	return returnedAlias, nil
}
