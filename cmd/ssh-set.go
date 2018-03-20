package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var SshSetCmd = &cobra.Command{
	Use:   "set <name> <path>",
	Short: "ConoHaに登録済みのSSH Keyと秘密鍵ファイルパスの紐付け",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := args[1]
		if !util.IsExist(path) {
			fmt.Println("存在しないファイルパスです。")
			return
		}

		try := util.Try
		try(conoha.SshAdd(name, path))
	},
}

var SshSetCmdHelp = `Usage:
  conoha ssh set <name> <path>
  conoha ssh set -h

Args:
  name: ConoHaに登録済みのSSH Key名
  path: 対応する秘密鍵のファイルパス (例: ~/.ssh/id_rsa)

Help:
  - ConoHaに登録済みのSSH Key名の一覧を表示
    $ conoha info ssh
`

func init() {
	SshCmd.AddCommand(SshSetCmd)
	SshSetCmd.SetUsageTemplate(SshSetCmdHelp)
}
