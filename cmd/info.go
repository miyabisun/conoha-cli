package cmd

import (
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "APIからプランやイメージ一覧情報を取得",
	Long:  "get Infomation from ConoHa API (require logged in).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(InfoCmd)
}
