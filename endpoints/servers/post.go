package servers

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

func Post(tenantId string, tokenId string, name string, image_id string, flavor_id string, sshKey string) error {
	res := &endpoints.Response{}
	url := fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/servers", tenantId)
	reqBody := fmt.Sprintf(`{"server": {"imageRef": "%s", "flavorRef": "%s", "metadata": {"instance_name_tag": "%s"}, "key_name": "%s"} }`, image_id, flavor_id, name, sshKey)
	return endpoints.Post(url, tokenId, reqBody, res)
}
