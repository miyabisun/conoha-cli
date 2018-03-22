package conoha

import (
	"errors"
	"fmt"

	"github.com/miyabisun/conoha-cli/util"
	"golang.org/x/crypto/ssh"
)

func SshRead(name string) (ssh.AuthMethod, error) {
	path := SshPath(name)
	if path == "" {
		return nil, errors.New(fmt.Sprintf("SSH Key <%s> は登録されていません", name))
	}

	return util.ReadPem(path)
}
