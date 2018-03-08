package sshClient

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/miyabisun/conoha-cli/config/status"
	"github.com/miyabisun/conoha-cli/util"
)

func OpenSSH() {
	try := util.Try

	sshInfo := &status.SshConfig{}
	try(status.Load(sshInfo))

	command := exec.Command("ssh", "-i", sshInfo.IdentityFile, fmt.Sprintf("%s@%s", sshInfo.User, sshInfo.HostName))
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Run()
}
