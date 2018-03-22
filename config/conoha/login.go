package conoha

import (
	"encoding/json"

	"github.com/miyabisun/conoha-cli/endpoints"
	"github.com/miyabisun/conoha-cli/endpoints/tokens"
)

type jsonAccess struct {
	Access jsonToken `json:"access"`
}
type jsonToken struct {
	Token ConfigToken `json:"token"`
}

func Login(auth *ConfigAuth, token *ConfigToken) error {
	var res endpoints.Response
	err := tokens.Post(auth.User, auth.Pass, auth.TenantId, &res)
	if err != nil {
		return err
	}

	err = parseToken(&res.Body, token)
	return err
}

func parseToken(body *[]byte, token *ConfigToken) error {
	var access jsonAccess
	err := json.Unmarshal(*body, &access)
	if err != nil {
		return err
	}
	token.Id = access.Access.Token.Id
	token.IssuedAt = access.Access.Token.IssuedAt
	token.Expires = access.Access.Token.Expires
	return nil
}
