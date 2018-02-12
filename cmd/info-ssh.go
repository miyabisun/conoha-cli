package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/keypairs"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var InfoSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "get registed ssh keypair-name ConoHa API.",
	Long:  "get registed ssh keypair-name API (require logged in).",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try
		try(conoha.Refresh())

		config := &conoha.Config{}
		try(conoha.Read(config))

		keypairs := &[]endpoint.Keypair{}
		try(endpoint.Get(config.Auth.TenantId, config.Token.Id, keypairs))

		for _, item := range *keypairs {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoSshCmd)
}
