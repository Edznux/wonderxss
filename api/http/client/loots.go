package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetLoots() ([]api.Loot, error) {
	var res api.Response
	var loots []api.Loot
	res, err := c.doAuthAPIRequest("GET", "/loots/", nil)
	if err != nil {
		return []api.Loot{}, errors.New("Couldn't get loots " + err.Error())
	}
	if res.Data != nil {
		for _, c := range res.Data.([]interface{}) {
			tmpLoot := api.Loot{}
			mapstructure.Decode(c, &tmpLoot)
			loots = append(loots, tmpLoot)
		}
	}

	return loots, nil
}

func (c *Client) GetLoot(id string) (api.Loot, error) {
	var res api.Response
	var loot api.Loot
	res, err := c.doAuthAPIRequest("GET", "/loots/"+id, nil)
	if err != nil {
		return api.Loot{}, errors.New("Couldn't get Loot " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &loot)
	}

	return loot, nil
}

func (c *Client) AddLoot(data string) (api.Loot, error) {
	return api.Loot{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteLoot(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/loots/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete loot " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
