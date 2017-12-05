package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/miyabisun/conoha-cli/conf"
)

func LoginFrom(auth *conf.ConfAuth) ([]byte, int, error) {
	client := &http.Client{}
	reqBody := fmt.Sprintf("{\"auth\":{\"passwordCredentials\":{\"username\":\"%s\",\"password\":\"%s\"},\"tenantId\":\"%s\"}}", auth.User, auth.Pass, auth.TenantId)
	buffer := bytes.NewBufferString(reqBody)
	request, err := http.NewRequest("POST", "https://identity.tyo1.conoha.io/v2.0/tokens", buffer)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return nil, statusCode, errors.New("login is failed.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, statusCode, err
}

func TokenId() (string, error) {
	config, err := conf.Read()
	if err != nil {
		return "", err
	}
	if isRequiredLogin(&config) {
		return "", errors.New("required login.")
	}

	if isRequiredRelogin(&config) {
		err := relogin()
		if err != nil {
			return "", err
		}
		config, _ = conf.Read()
	}
	return config.Token.Id, nil
}

func ToToken(body []byte) (conf.ConfToken, error) {
	access := conf.JsonAccess{}
	err := json.Unmarshal(body, &access)
	if err != nil {
		return conf.ConfToken{}, err
	}
	return access.Access.Token, err
}

func relogin() error {
	config, err := conf.Read()
	if err != nil {
		return err
	}

	// login
	body, statusCode, err := LoginFrom(&config.Auth)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return errors.New("status code isnt 200: " + string(body))
	}

	// update config
	token, err := ToToken(body)
	if err != nil {
		return err
	}
	config.Token = token
	return conf.Write(&config)
}

func isRequiredLogin(config *conf.Conf) bool {
	return (config.Token.Id == "" || config.Token.Expires == "" || config.Auth.User == "" || config.Auth.Pass == "" || config.Auth.TenantId == "")
}

func isRequiredRelogin(config *conf.Conf) bool {
	format := "2006-01-02T15:04:05-0700"
	expires, err := time.Parse(format, config.Token.Expires)
	return (err != nil || expires.Unix() < time.Now().Unix())
}
