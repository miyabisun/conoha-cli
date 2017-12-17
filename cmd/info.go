package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "get Infomation from ConoHa API.",
	Long:  "get Infomation from ConoHa API (require logged in).",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("require at one arg")
			return
		}

		var items []util.Item
		err := util.Info(args[0], &items)
		if err != nil {
			panic(err)
		}

		for _, item := range items {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(InfoCmd)
}

func findToken() string {
	return ""
}
