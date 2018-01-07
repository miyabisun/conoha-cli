package keypairs

import (
	"encoding/json"
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

type Keypair struct {
	Name string
}
type jsonKeypairs struct {
	Items []jsonKeypair `json:"keypairs"`
}
type jsonKeypair struct {
	Keypair Keypair
}

func Get(tenantId string, tokenId string, items *[]Keypair) error {
	var res endpoints.Response
	url := fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/os-keypairs", tenantId)
	err := endpoints.Get(url, tokenId, &res)
	if err != nil {
		return err
	}

	var parsed jsonKeypairs
	err = json.Unmarshal(res.Body, &parsed)
	if err != nil {
		return err
	}
	for _, item := range parsed.Items {
		*items = append(*items, item.Keypair)
	}

	return nil
}
