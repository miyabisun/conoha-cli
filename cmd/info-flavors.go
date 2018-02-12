package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/flavors"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var InfoFlavorsCmd = &cobra.Command{
	Use:   "flavors",
	Short: "get flavors ConoHa API.",
	Long:  "get flavors ConoHa API (require logged in).",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try

		try(conoha.Refresh())

		config := &conoha.Config{}
		try(conoha.Read(config))
		tenantId := config.Auth.TenantId
		tokenId := config.Token.Id

		flavors := &[]endpoint.Flavor{}
		try(endpoint.Get(tenantId, tokenId, flavors))

		for _, item := range *flavors {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoFlavorsCmd)
}
