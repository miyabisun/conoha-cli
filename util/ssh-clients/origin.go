package sshClient

import (
	"fmt"
	"os"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/status"
	"github.com/miyabisun/conoha-cli/util"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func Origin() error {
	try := util.Try

	sshInfo := &status.SshConfig{}
	try(status.Load(sshInfo))

	key, err := conoha.SshRead(sshInfo.KeyName)
	try(err)

	sshConfig := &ssh.ClientConfig{
		User:            "root",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{key},
	}
	host := fmt.Sprintf("%s:%s", sshInfo.HostName, "22")

	conn, err := ssh.Dial("tcp", host, sshConfig)
	try(err)
	defer conn.Close()

	session, err := conn.NewSession()
	try(err)
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	fs := int(os.Stdin.Fd())
	if terminal.IsTerminal(fs) {
		state, err := terminal.MakeRaw(fs)
		try(err)
		defer terminal.Restore(fs, state)
		termWidth, termHeight, err := terminal.GetSize(fs)
		try(err)
		try(session.RequestPty("xterm-256color", termHeight, termWidth, modes))
	}

	try(session.Shell())
	try(session.Wait())
	return nil
}
