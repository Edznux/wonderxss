package client

import (
	"errors"

	"github.com/edznux/wonderxss/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func (c *Client) GetUserByName(name string) (api.User, error) {
	return api.User{}, errors.New("Not implemented yet")
}

func (c *Client) GetUser(id string) (api.User, error) {
	var res api.Response
	var user api.User
	res, err := c.doAuthAPIRequest("GET", "/users/"+id, nil)
	if err != nil {
		return api.User{}, errors.New("Couldn't get User " + id + ": " + err.Error())
	}

	if res.Data != nil {
		mapstructure.Decode(res.Data, &user)
	}

	return user, nil
}

func (c *Client) CreateOTP(userID string, secret string, otp string) (api.User, error) {
	return api.User{}, errors.New("Not implemented yet")
}

func (c *Client) CreateUser(username, password string) (api.User, error) {
	return api.User{}, errors.New("Not implemented yet")
}

func (c *Client) DeleteUser(id string) error {
	res, err := c.doAuthAPIRequest("DELETE", "/users/"+id, nil)
	if err != nil {
		return errors.New("Couldn't delete users " + id + " " + err.Error())
	}
	log.Info(res)
	return err
}
