package images

import (
	"encoding/json"

	"github.com/miyabisun/conoha-cli/endpoints"
)

type Image struct {
	Id   string
	Name string
}

type jsonImages struct {
	Items []Image `json:"images"`
}

func Get(tokenId string, items *[]Image) error {
	var res endpoints.Response
	url := "https://image-service.tyo1.conoha.io/v2/images"
	err := endpoints.Get(url, tokenId, &res)
	if err != nil {
		return err
	}

	var parsed jsonImages
	err = json.Unmarshal(res.Body, &parsed)
	if err != nil {
		return err
	}
	*items = parsed.Items

	return nil
}
