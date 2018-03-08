package sshClient

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/miyabisun/conoha-cli/config/status"
	"github.com/miyabisun/conoha-cli/util"
)

func Mosh() {
	try := util.Try

	sshInfo := &status.SshConfig{}
	try(status.Load(sshInfo))

	cmdStr := fmt.Sprintf("mosh --ssh=\"ssh -i %s\" %s@%s", sshInfo.IdentityFile, sshInfo.User, sshInfo.HostName)
	command := exec.Command("bash", "-c", cmdStr)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Run()
}
