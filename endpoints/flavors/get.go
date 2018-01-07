package flavors

import (
	"encoding/json"
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

type Flavor struct {
	Id   string
	Name string
}

type jsonFlavors struct {
	Items []Flavor `json:"flavors"`
}

func Get(tenantId string, tokenId string, items *[]Flavor) error {
	var res endpoints.Response
	url := fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/flavors", tenantId)
	err := endpoints.Get(url, tokenId, &res)
	if err != nil {
		return err
	}

	var parsed jsonFlavors
	err = json.Unmarshal(res.Body, &parsed)
	if err != nil {
		return err
	}
	*items = parsed.Items

	return nil
}
