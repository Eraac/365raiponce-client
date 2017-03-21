package raiponce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	authURI         = "/oauth/v2/token"
	formatTokenFile = "token-%s.json"
)

// Token store token used for authentication
type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// Authentication is used for authenticate the client
type authentication struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (client *Client) Login(username string, password string) {
	auth := authentication{
		ClientID:     client.config.ClientID,
		ClientSecret: client.config.ClientSecret,
	}

	token, exist := getTokenViaFile(username)

	if exist && accessTokenIsExpire(token) {
		client.authByRefresh(auth, token, username, password)
	} else if !exist {
		client.authByPassword(auth, username, password)
	} else {
		client.token = token
	}
}

func (client *Client) auth(auth authentication, username string) error {
	token := token{}

	err := client.post(authURI, auth, &token)

	client.token = token

	updateAuthFile(username, token)

	return err
}

func (client *Client) authByPassword(auth authentication, username string, password string) error {
	auth.GrantType = grantTypePassword
	auth.Username = username
	auth.Password = password

	return client.auth(auth, username)
}

func (client *Client) authByRefresh(auth authentication, token token, username string, password string) error {
	auth.GrantType = grantTypeRefreshToken
	auth.RefreshToken = token.RefreshToken

	err := client.auth(auth, username)

	if err != nil {
		err = client.authByPassword(auth, username, password)
	}

	return err
}

func getTokenViaFile(username string) (token, bool) {
	var token token

	file, err := os.Open(fmt.Sprintf(formatTokenFile, username))

	if err != nil {
		return token, false
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&token)

	if err != nil {
		log.Fatalln(err)
	}

	return token, true
}

func accessTokenIsExpire(token token) bool {
	timestamp := time.Now().Unix()

	return timestamp > int64(token.ExpiresIn)
}

func updateAuthFile(username string, token token) {
	token.ExpiresIn = int(time.Now().Unix()) + token.ExpiresIn

	output, err := json.Marshal(token)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf(formatTokenFile, username), output, 0600)

	if err != nil {
		log.Fatal(err)
	}
}
