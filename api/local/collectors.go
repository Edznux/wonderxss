package local

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

// GetCollectors return all the triggered payload stored in the database.
func (local *Local) GetCollectors() ([]api.Collector, error) {
	fmt.Println("api.GetCollectors")
	data, err := local.store.GetCollectors()
	if err != nil {
		return nil, err
	}
	collectors := []api.Collector{}
	fmt.Println("Collectors from store: ", data)
	for _, p := range data {
		tmp := api.Collector{}
		collectors = append(collectors, tmp.FromStorage(p))
	}

	return collectors, nil
}

func (local *Local) GetCollector(id string) (api.Collector, error) {
	collector, err := local.store.GetCollector(id)
	if err != nil {
		return api.Collector{}, err
	}
	res := api.Collector{}
	return res.FromStorage(collector), nil
}

func (local *Local) AddCollector(data string) (api.Collector, error) {
	var returnedCollector api.Collector
	fmt.Printf("AddCollector(\"%s\")\n", data)
	c := models.Collector{
		ID:        uuid.New().String(),
		Data:      data,
		CreatedAt: time.Now(),
	}
	fmt.Println(c)
	collector, err := local.store.CreateCollector(c)
	if err != nil {
		return api.Collector{}, err
	}

	return returnedCollector.FromStorage(collector), nil
}

func (local *Local) DeleteCollector(id string) error {
	e := models.Collector{ID: id}
	err := local.store.DeleteCollector(e)
	if err != nil {
		return err
	}
	return nil
}
