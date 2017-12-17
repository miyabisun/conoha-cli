package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/miyabisun/conoha-cli/conf"
)

type endpoint struct {
	method string
	url    string
	tenant bool
}

type jsonImages struct {
	Items []Item `json:"images"`
}

type jsonflavors struct {
	Items []Item `json:"flavors"`
}

type Item struct {
	Id   string
	Name string
}

var endpoints = map[string]endpoint{
	"images": {
		method: "GET",
		url:    "https://image-service.tyo1.conoha.io/v2/images",
		tenant: false,
	},
	"flavors": {
		method: "GET",
		url:    "https://compute.tyo1.conoha.io/v2/%s/flavors",
		tenant: true,
	},
}

func Info(target string, items *[]Item) error {
	tokenId, err := TokenId()
	if err != nil {
		return err
	}

	config, _ := conf.Read()
	endpoint := endpoints[target]
	var url string
	if endpoint.tenant {
		url = fmt.Sprintf(endpoint.url, config.Auth.TenantId)
	} else {
		url = endpoint.url
	}
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-Auth-Token", tokenId)
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO:ここ綺麗に書きたい
	if target == "images" {
		var parsed jsonImages
		err := json.Unmarshal(body, &parsed)
		if err != nil {
			return err
		}
		*items = parsed.Items
	} else if target == "flavors" {
		var parsed jsonflavors
		err := json.Unmarshal(body, &parsed)
		if err != nil {
			return err
		}
		*items = parsed.Items
	}
	return nil
}
