package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/flavors"
	"github.com/spf13/cobra"
)

var InfoFlavorsCmd = &cobra.Command{
	Use:   "flavors",
	Short: "get flavors ConoHa API.",
	Long:  "get flavors ConoHa API (require logged in).",
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

		flavors := &[]endpoint.Flavor{}
		err = endpoint.Get(tenantId, tokenId, flavors)
		if err != nil {
			panic(err)
		}

		for _, item := range *flavors {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoFlavorsCmd)
}
