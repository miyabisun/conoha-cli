package status

import (
	"errors"
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
)

func Load(config *SshConfig) error {
	st := &Config{}
	if err := Read(st); err != nil {
		return err
	}

	config.Name = st.Name
	config.HostName = st.IpAddr
	config.User = "root"
	config.KeyName = st.KeyName
	config.IdentityFile = "~/.ssh/id_rsa"

	path := conoha.SshPath(st.KeyName)
	if path == "" {
		return errors.New(fmt.Sprintf("SSH Key <%s> は登録されていません", st.KeyName))
	}
	config.IdentityFile = path

	return nil
}
