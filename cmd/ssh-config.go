package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/status"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var name string

func init() {
	RootCmd.AddCommand(SshConfigCmd)
	SshConfigCmd.Flags().StringVarP(&name, "name", "n", "", "Hostname")
}

var SshConfigCmd = &cobra.Command{
	Use:   "ssh-config",
	Short: "Show SSH config.",
	Long:  "Show SSH config.",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try
		config := &conoha.Config{}
		try(conoha.Read(config))

		sshConf := &status.SshConfig{}
		try(status.Load(sshConf))

		format := `Host %s
  HostName %s
  User %s
  IdentityFile %s`
		fmt.Printf(format, sshConf.Name, sshConf.HostName, sshConf.User, sshConf.IdentityFile)
	},
}
