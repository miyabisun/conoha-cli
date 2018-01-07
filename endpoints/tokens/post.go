package tokens

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

func Post(user string, pass string, tenantId string, res *endpoints.Response) error {
	url := "https://identity.tyo1.conoha.io/v2.0/tokens"
	body := fmt.Sprintf("{\"auth\":{\"passwordCredentials\":{\"username\":\"%s\",\"password\":\"%s\"},\"tenantId\":\"%s\"}}", user, pass, tenantId)
	return endpoints.Post(url, "", body, res)
}
