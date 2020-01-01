package api

import (
	"fmt"
	"log"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

func GetInjections() ([]Injection, error) {
	data, err := store.GetInjections()
	if err != nil {
		return nil, err
	}
	injections := []Injection{}
	for _, p := range data {
		tmp := Injection{}
		injections = append(injections, tmp.fromStorage(p))
	}

	return injections, nil
}

func GetInjection(id string) (Injection, error) {
	injection, err := store.GetInjection(id)
	if err != nil {
		return Injection{}, err
	}
	res := Injection{}
	return res.fromStorage(injection), nil
}

func AddInjection(name string, content string) (models.Injection, error) {
	fmt.Printf("AddInjection(\"%s\", \"%s\")\n", name, content)
	p := models.Injection{
		ID:      uuid.New().String(),
		Name:    name,
		Content: content,
	}
	returnedInjection, err := store.CreateInjection(p)
	if err != nil {
		log.Println("could not add Injection:", err)
		return models.Injection{}, err
	}

	return returnedInjection, nil
}

func DeleteInjection(id string) error {
	e := models.Injection{ID: id}
	err := store.DeleteInjection(e)
	if err != nil {
		return err
	}
	return nil
}
