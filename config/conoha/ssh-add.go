package conoha

import (
	"errors"
	"fmt"

	endpoint "github.com/miyabisun/conoha-cli/endpoints/keypairs"
	"github.com/miyabisun/conoha-cli/util"
	"golang.org/x/crypto/ssh"
)

func SshAdd(name string, path string) error {
	config := &Config{}
	err := Read(config)
	if err != nil {
		return err
	}

	keypair := &endpoint.Keypair{}
	err = endpoint.Show(config.Auth.TenantId, config.Token.Id, name, keypair)
	if err != nil {
		return err
	}
	if keypair.Name == "" {
		return errors.New(fmt.Sprintf("SSH Key <%s> はConoHaに登録されていません", name))
	}

	key, err := util.ReadRsaPrivateKey(path)
	if err != nil {
		return err
	}
	_, err = ssh.NewSignerFromKey(key)
	if err != nil {
		return err
	}
	_, err = util.ReadPem(path)
	if err != nil {
		return err
	}

	ssh := &ConfigSsh{}
	for _, it := range config.Ssh {
		if it.Name == name {
			*ssh = it
			break
		}
	}
	if ssh.Name != "" {
		ssh.Path = path
	} else {
		ssh.Name = name
		ssh.Path = path
		config.Ssh = append(config.Ssh, *ssh)
	}

	return Write(config)
}
