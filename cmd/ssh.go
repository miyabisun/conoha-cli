package cmd

import (
	"os/exec"

	sshClient "github.com/miyabisun/conoha-cli/util/ssh-clients"
	"github.com/spf13/cobra"
)

var isMosh bool

func init() {
	RootCmd.AddCommand(SshCmd)
	SshCmd.Flags().BoolVarP(&isMosh, "mosh", "m", false, "Use Mosh (mobile shell)")
}

var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "VPSインスタンスへのSSH接続",
	Run: func(cmd *cobra.Command, args []string) {
		if isMosh {
			sshClient.Mosh()
		} else if exec.Command("ssh", "-V").Run() == nil {
			sshClient.OpenSSH()
		} else {
			sshClient.Origin()
		}
	},
}
