package raiponce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	grantTypePassword     = "password"
	grantTypeRefreshToken = "refresh_token"
)

type ClientConfig struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	APIVersion   string
	Locale       string
}

// Client is used for request the api
type Client struct {
	config ClientConfig
	token  token
	client *http.Client
}

// New create a new Client for use API
func NewClient(config ClientConfig) *Client {
	return &Client{
		config: config,
		client: &http.Client{},
	}
}

func (client *Client) cget(uri string, response interface{}, query *QueryFilter) error {
	if query != nil {
		uri += fmt.Sprintf("?%s", query.buildQuery())
	}

	return client.sendRequest(http.MethodGet, uri, nil, response)
}

func (client *Client) get(uri string, response interface{}) error {
	return client.sendRequest(http.MethodGet, uri, nil, response)
}

func (client *Client) post(uri string, body interface{}, response interface{}) error {
	return client.sendRequest(http.MethodPost, uri, body, response)
}

func (client *Client) patch(uri string, body interface{}, response interface{}) error {
	return client.sendRequest(http.MethodPatch, uri, body, response)
}

func (client *Client) remove(uri string) error {
	return client.sendRequest(http.MethodDelete, uri, nil, nil)
}

func (client *Client) addHeader(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Accept-Version", client.config.APIVersion)
	request.Header.Add("Accept-Language", client.config.Locale)

	if len(client.token.AccessToken) > 0 {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token.AccessToken))
	}
}

func (client *Client) sendRequest(method string, uri string, body interface{}, response interface{}) error {
	var request *http.Request
	var err error

	url := client.config.BaseURL + uri

	if body != nil {
		jsonBody := structToReader(body)

		request, err = http.NewRequest(method, url, jsonBody)
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Fatalln(err)
	}

	client.addHeader(request)

	res, err := client.client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	// TODO handle error

	if response != nil {
		json.NewDecoder(res.Body).Decode(&response)
	}

	if res.StatusCode > 300 {
		return errors.New("")
	}

	return nil
}

func structToReader(s interface{}) *bytes.Reader {
	sliceBytes, err := json.Marshal(s)

	if err != nil {
		log.Fatalln(err)
	}

	reader := bytes.NewReader(sliceBytes)

	return reader
}
