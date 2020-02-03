package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetInjections() ([]api.Injection, error) {
	var res api.Response
	var injections []api.Injection
	res, err := c.doAuthAPIRequest("GET", "/injections", nil)
	if err != nil {
		return []api.Injection{}, errors.New("Couldn't get Injections: " + err.Error())
	}
	if res.Data != nil {
		for _, i := range res.Data.([]interface{}) {
			tmpInjection := api.Injection{}
			mapstructure.Decode(i, &tmpInjection)
			injections = append(injections, tmpInjection)
		}
	}
	return injections, nil
}

func (c *Client) GetInjection(id string) (api.Injection, error) {
	var res api.Response
	var injection api.Injection
	res, err := c.doAuthAPIRequest("GET", "/injections/"+id, nil)
	if err != nil {
		return api.Injection{}, errors.New("Couldn't get Injection " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &injection)
	}

	return injection, nil
}

func (c *Client) AddInjection(name string, content string) (api.Injection, error) {
	return api.Injection{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteInjection(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/injections/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete injection " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
