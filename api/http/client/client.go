package client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/edznux/wonderxss/api"
)

type Client struct {
	Version   string
	Host      string
	Protocol  string
	Port      int
	jwtToken  string
	apiPrefix string
}

func New() *Client {
	c := Client{}
	c.Version = "v1"
	c.apiPrefix = "/api/" + c.Version
	c.Protocol = "http://"
	c.Host = "localhost"
	c.Port = 80
	return &c
}

func (c *Client) formatURLApi(path string) string {
	return c.Protocol + c.Host + ":" + strconv.Itoa(c.Port) + path
}

func (c *Client) doRequest(method string, path string, body io.Reader) (api.Response, error) {
	log.Println(method, path)
	var result api.Response
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, c.formatURLApi(path), body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	req.Header.Add("Authorization", "Bearer "+c.jwtToken)

	response, err := netClient.Do(req)
	err = json.NewDecoder(response.Body).Decode(&result)
	return result, err
}

func (c *Client) doAuthRequest(method string, path string, body io.Reader) (api.Response, error) {
	var result api.Response
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	req.Header.Add("Authorization", "Bearer "+c.jwtToken)

	response, err := netClient.Do(req)
	err = json.NewDecoder(response.Body).Decode(&result)
	return result, err
}

func (c *Client) doAPIRequest(method string, path string, body io.Reader) (api.Response, error) {
	return c.doRequest(method, c.apiPrefix+path, body)
}

func (c *Client) doAuthAPIRequest(method string, path string, body io.Reader) (api.Response, error) {
	return c.doAuthRequest(method, c.apiPrefix+path, body)
}

func (c *Client) Login(user, password, otp string) (string, error) {
	var result api.Response
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	form := url.Values{}
	form.Add("login", user)
	form.Add("password", password)
	form.Add("token", otp)

	req, err := http.NewRequest("POST", c.formatURLApi("/login"), strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	log.Println(req)
	response, err := netClient.Do(req)
	err = json.NewDecoder(response.Body).Decode(&result)
	if result.Error != "" {
		log.Println(err)
		return "", errors.New(result.Error)
	}
	token := result.Data.(string)
	if token == "" {
		log.Println("Empty token")
		return "", errors.New("Could not connect")
	}
	return token, nil
}

func (c *Client) GetHealth() (api.Response, error) {
	res, err := c.doAPIRequest("GET", "/healthz", nil)
	if err != nil {
		return res, errors.New("Couldn't get HEALTHZ informations: " + err.Error())
	}
	return res, nil
}
