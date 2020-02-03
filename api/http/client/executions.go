package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetExecutions() ([]api.Execution, error) {
	var res api.Response
	var executions []api.Execution
	res, err := c.doAuthAPIRequest("GET", "/executions/", nil)
	if err != nil {
		return []api.Execution{}, errors.New("Couldn't get executions " + err.Error())
	}
	if res.Data != nil {
		for _, e := range res.Data.([]interface{}) {
			tmpExecutions := api.Execution{}
			mapstructure.Decode(e, &tmpExecutions)
			executions = append(executions, tmpExecutions)
		}
	}

	return executions, nil
}

func (c *Client) GetExecution(id string) (api.Execution, error) {
	var res api.Response
	var execution api.Execution
	res, err := c.doAuthAPIRequest("GET", "/executions/"+id, nil)
	if err != nil {
		return api.Execution{}, errors.New("Couldn't get Execution " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &execution)
	}

	return execution, nil
}

func (c *Client) AddExecution(payloadID string, aliasID string) (api.Execution, error) {
	return api.Execution{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteExecution(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/executions/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete execution " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
