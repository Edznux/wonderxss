package api

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

// GetExecutions return all the triggered payload stored in the database.
func GetExecutions() ([]Execution, error) {
	fmt.Println("api.GetExecutions")
	data, err := store.GetExecutions()
	if err != nil {
		return nil, err
	}
	executions := []Execution{}
	fmt.Println("Executions from store: ", data)
	for _, p := range data {
		tmp := Execution{}
		executions = append(executions, tmp.fromStorage(p))
	}

	return executions, nil
}

func GetExecution(id string) (Execution, error) {
	execution, err := store.GetExecution(id)
	if err != nil {
		return Execution{}, err
	}
	res := Execution{}
	return res.fromStorage(execution), nil
}

func AddExecution(payloadID string, aliasID string) (models.Execution, error) {
	fmt.Printf("AddExecution(\"%s\", \"%s\")\n", payloadID, aliasID)
	l := models.Execution{
		ID:          uuid.New().String(),
		PayloadID:   payloadID,
		AliasID:     aliasID,
		TriggeredAt: time.Now(),
	}
	fmt.Println(l)
	returnedAlias, err := store.CreateExecution(l, aliasID)
	if err != nil {
		return models.Execution{}, err
	}

	return returnedAlias, nil
}
