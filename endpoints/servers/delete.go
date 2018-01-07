package servers

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

func Delete(tenantId string, tokenId string, id string) error {
	res := &endpoints.Response{}
	url := fmt.Sprintf(`https://compute.tyo1.conoha.io/v2/%s/servers/%s`, tenantId, id)
	return endpoints.Delete(url, tokenId, res)
}
