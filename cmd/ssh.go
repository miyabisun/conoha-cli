package cmd

import (
	"fmt"
	"os"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/status"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "VPSインスタンスにSSH接続します",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try

		config := &conoha.Config{}
		try(conoha.Read(config))

		statusConf := &status.Config{}
		try(status.Read(statusConf))

		key, err := conoha.SshRead(statusConf.KeyName)
		try(err)

		sshConfig := &ssh.ClientConfig{
			User:            "root",
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth:            []ssh.AuthMethod{key},
		}
		host := fmt.Sprintf("%s:%s", statusConf.IpAddr, "22")

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
			ssh.ECHO:          0,
			ssh.TTY_OP_ISPEED: 14400,
			ssh.TTY_OP_OSPEED: 14400,
		}
		term := os.Getenv("TERM")
		try(session.RequestPty(term, 25, 80, modes))
		try(session.Shell())
		try(session.Wait())
	},
}

func init() {
	RootCmd.AddCommand(SshCmd)
}
