package api

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

// GetCollectors return all the triggered payload stored in the database.
func GetCollectors() ([]Collector, error) {
	fmt.Println("api.GetCollectors")
	data, err := store.GetCollectors()
	if err != nil {
		return nil, err
	}
	collectors := []Collector{}
	fmt.Println("Collectors from store: ", data)
	for _, p := range data {
		tmp := Collector{}
		collectors = append(collectors, tmp.fromStorage(p))
	}

	return collectors, nil
}

func GetCollector(id string) (Collector, error) {
	collector, err := store.GetCollector(id)
	if err != nil {
		return Collector{}, err
	}
	res := Collector{}
	return res.fromStorage(collector), nil
}

func AddCollector(payloadID string, data string) (models.Collector, error) {
	fmt.Printf("AddCollector(\"%s\", \"%s\")\n", payloadID, data)
	c := models.Collector{
		ID:        uuid.New().String(),
		PayloadID: payloadID,
		Data:      data,
		CreatedAt: time.Now(),
	}
	fmt.Println(c)
	returnedAlias, err := store.CreateCollector(c)
	if err != nil {
		return models.Collector{}, err
	}

	return returnedAlias, nil
}

func DeleteCollector(id string) error {
	e := models.Collector{ID: id}
	err := store.DeleteCollector(e)
	if err != nil {
		return err
	}
	return nil
}
