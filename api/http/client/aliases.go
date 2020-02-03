package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetAliases() ([]api.Alias, error) {
	var res api.Response
	aliases := []api.Alias{}
	res, err := c.doAuthAPIRequest("GET", "/aliases", nil)
	if err != nil {
		return []api.Alias{}, errors.New("Couldn't get Aliases: " + err.Error())
	}
	if res.Data != nil {
		for _, a := range res.Data.([]interface{}) {
			tmpAlias := api.Alias{}
			mapstructure.Decode(a, &tmpAlias)
			aliases = append(aliases, tmpAlias)
		}
	}
	return aliases, nil
}

func (c *Client) GetAlias(id string) (api.Alias, error) {
	var res api.Response
	var alias api.Alias
	res, err := c.doAuthAPIRequest("GET", "/aliases/"+id, nil)
	if err != nil {
		return api.Alias{}, errors.New("Couldn't get Alias " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &alias)
	}

	return alias, nil
}

func (c *Client) GetAliasByID(id string) (api.Alias, error) {
	return c.GetAlias(id)
}

func (c *Client) GetAliasByPayloadID(id string) (api.Alias, error) {
	return api.Alias{}, errors.New("Not implemented yet")
}

func (c *Client) AddAlias(name string, payloadId string) (api.Alias, error) {
	return api.Alias{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteAlias(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/aliases/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete alias " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
