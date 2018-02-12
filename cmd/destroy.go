package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/servers"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DestroyCmd)
}

var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy in ConoHa API.",
	Long:  "destroy in ConoHa API(required logged in)",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try
		try(conoha.Refresh())

		confSpec := &spec.Config{}
		try(spec.Read(confSpec))
		name := confSpec.Name

		config := &conoha.Config{}
		try(conoha.Read(config))
		tenantId := config.Auth.TenantId
		tokenId := config.Token.Id

		servers := &[]endpoint.Server{}
		try(endpoint.Get(tenantId, tokenId, servers))

		var id string
		for _, it := range *servers {
			if it.Metadata.Instance_name_tag == name {
				id = it.Id
			}
		}
		if id == "" {
			fmt.Println("not found server.")
			return
		}

		try(endpoint.Delete(tenantId, tokenId, id))
		fmt.Println("delete succesful.")
	},
}
