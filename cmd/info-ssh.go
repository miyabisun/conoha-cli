package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/keypairs"
	"github.com/spf13/cobra"
)

var InfoSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "get registed ssh keypair-name ConoHa API.",
	Long:  "get registed ssh keypair-name API (require logged in).",
	Run: func(cmd *cobra.Command, args []string) {
		err := conoha.Refresh()
		if err != nil {
			panic(err)
		}

		config := &conoha.Config{}
		err = conoha.Read(config)
		if err != nil {
			panic(err)
		}
		tenantId := config.Auth.TenantId
		tokenId := config.Token.Id

		keypairs := &[]endpoint.Keypair{}
		err = endpoint.Get(tenantId, tokenId, keypairs)
		if err != nil {
			panic(err)
		}

		for _, item := range *keypairs {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoSshCmd)
}
