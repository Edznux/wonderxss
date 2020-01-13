package local

import (
	"fmt"
	"time"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
)

// GetExecutions return all the triggered payload stored in the database.
func (local *Local) GetExecutions() ([]api.Execution, error) {
	fmt.Println("api.GetExecutions")
	data, err := local.store.GetExecutions()
	if err != nil {
		return nil, err
	}
	executions := []api.Execution{}
	fmt.Println("Executions from store: ", data)
	for _, p := range data {
		tmp := api.Execution{}
		executions = append(executions, tmp.FromStorage(p))
	}

	return executions, nil
}

func (local *Local) GetExecution(id string) (api.Execution, error) {
	execution, err := local.store.GetExecution(id)
	if err != nil {
		return api.Execution{}, err
	}
	res := api.Execution{}
	return res.FromStorage(execution), nil
}

func (local *Local) AddExecution(payloadID string, aliasID string) (api.Execution, error) {
	var returnedExecution api.Execution
	fmt.Printf("AddExecution(\"%s\", \"%s\")\n", payloadID, aliasID)
	l := models.Execution{
		ID:          uuid.New().String(),
		PayloadID:   payloadID,
		AliasID:     aliasID,
		TriggeredAt: time.Now(),
	}
	fmt.Println(l)
	execution, err := local.store.CreateExecution(l, aliasID)
	if err != nil {
		return returnedExecution, err
	}

	return returnedExecution.FromStorage(execution), nil
}

func (local *Local) DeleteExecution(id string) error {
	e := models.Execution{ID: id}
	err := local.store.DeleteExecution(e)
	if err != nil {
		return err
	}
	return nil
}
