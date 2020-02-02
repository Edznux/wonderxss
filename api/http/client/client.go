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
	"github.com/mitchellh/mapstructure"
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
	return &c
}

func (c *Client) formatURLApi(path string) string {
	return c.Protocol + c.Host + ":" + strconv.Itoa(c.Port) + path
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

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Warnln(err)
		return result, err
	}

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
	return errors.New("Not implemented yet")
}

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
	return errors.New("Not implemented yet")
}

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
	return errors.New("Not implemented yet")
}

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
	return errors.New("Not implemented yet")
}

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
	return errors.New("Not implemented yet")
}

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
	return errors.New("Not implemented yet")
}
