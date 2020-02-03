package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/config"
)

type Client struct {
	Version   string
	Host      string
	Protocol  string
	Port      int
	jwtToken  string
	apiPrefix string
}

func New(cfg config.Client) *Client {
	c := Client{}
	c.Version = cfg.Version
	c.apiPrefix = "/api/" + c.Version
	c.Protocol = "http://"
	c.Host = cfg.Host
	c.Port = cfg.Port
	c.jwtToken = cfg.Token
	log.Debugf("Version: %s, prefix: %s, protocol:%s, host:%s, port:%d", c.Version, c.apiPrefix, c.Protocol, c.Host, c.Port)
	return &c
}

func (c *Client) formatURLApi(path string) string {
	return c.Protocol + c.Host + ":" + strconv.Itoa(c.Port) + path
}
func (c *Client) setUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", "WonderXSS "+c.Version+" (https://github.com/edznux/wonderxss)")
}

func (c *Client) doRequest(method string, path string, body io.Reader) (api.Response, error) {
	log.Debugln(method, c.formatURLApi(path))
	var result api.Response
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, c.formatURLApi(path), body)
	if err != nil {
		log.Warnln(err)
		return result, err
	}
	c.setUserAgent(req)
	req.Header.Add("Authorization", "Bearer "+c.jwtToken)

	response, err := netClient.Do(req)
	if err != nil {
		log.Warnln(err)
		return result, err
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	return result, err
}

func (c *Client) doAuthRequest(method string, path string, body io.Reader) (api.Response, error) {
	var result api.Response
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(method, c.formatURLApi(path), body)
	if err != nil {
		log.Warnln(err)
		return result, err
	}
	c.setUserAgent(req)
	req.Header.Add("Authorization", "Bearer "+c.jwtToken)

	response, err := netClient.Do(req)
	if err != nil {
		log.Errorln(err)
		return result, err
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return api.Response{}, err
	}

	if result.Error != "" {
		return api.Response{}, errors.New(result.Error)
	}

	return result, nil
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
		log.Warnln(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := netClient.Do(req)
	err = json.NewDecoder(response.Body).Decode(&result)
	if result.Error != "" {
		log.Warnln(err)
		return "", errors.New(result.Error)
	}
	token := result.Data.(string)
	if token == "" {
		log.Warnln("Empty token")
		return "", errors.New("Could not connect")
	}
	return token, nil
}

func (c *Client) GetHealth() (string, error) {
	var res api.Response
	res, err := c.doAuthAPIRequest("GET", "/healthz", nil)
	if err != nil {
		return "", errors.New("Couldn't get HEALTHZ informations: " + err.Error())
	}
	return res.Data.(string), nil
}
