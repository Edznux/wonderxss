package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetPayloads() ([]api.Payload, error) {
	var res api.Response
	var payloads []api.Payload
	res, err := c.doAuthAPIRequest("GET", "/payloads", nil)
	if err != nil {
		return []api.Payload{}, errors.New("Couldn't get Payload: " + err.Error())
	}
	if res.Data != nil {
		for _, p := range res.Data.([]interface{}) {
			tmpPayload := api.Payload{}
			mapstructure.Decode(p, &tmpPayload)
			payloads = append(payloads, tmpPayload)
		}
	}
	return payloads, nil
}

func (c *Client) ServePayload(idOrAlias string) (string, error) {
	return "", errors.New("Not implemented yet")
}

func (c *Client) GetPayload(id string) (api.Payload, error) {
	var res api.Response
	var payload api.Payload
	res, err := c.doAuthAPIRequest("GET", "/payloads/"+id, nil)
	if err != nil {
		return api.Payload{}, errors.New("Couldn't get Payload " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &payload)
	}

	return payload, nil
}

func (c *Client) AddPayload(name string, content string, contentType string) (api.Payload, error) {
	return api.Payload{}, errors.New("Not implemented yet")
}

func (c *Client) DeletePayload(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/payloads/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete payload " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
