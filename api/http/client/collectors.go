package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetCollectors() ([]api.Collector, error) {
	var res api.Response
	var collectors []api.Collector
	res, err := c.doAuthAPIRequest("GET", "/collectors/", nil)
	if err != nil {
		return []api.Collector{}, errors.New("Couldn't get collectors " + err.Error())
	}
	if res.Data != nil {
		for _, c := range res.Data.([]interface{}) {
			tmpCollector := api.Collector{}
			mapstructure.Decode(c, &tmpCollector)
			collectors = append(collectors, tmpCollector)
		}
	}

	return collectors, nil
}

func (c *Client) GetCollector(id string) (api.Collector, error) {
	var res api.Response
	var collector api.Collector
	res, err := c.doAuthAPIRequest("GET", "/collectors/"+id, nil)
	if err != nil {
		return api.Collector{}, errors.New("Couldn't get Collector " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &collector)
	}

	return collector, nil
}

func (c *Client) AddCollector(data string) (api.Collector, error) {
	return api.Collector{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteCollector(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/collectors/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete collector " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
