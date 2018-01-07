package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/servers"
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
		err := conoha.Refresh()
		if err != nil {
			panic(err)
		}

		confSpec := &spec.Config{}
		err = spec.Read(confSpec)
		if err != nil {
			panic(err)
		}
		name := confSpec.Name

		config := &conoha.Config{}
		err = conoha.Read(config)
		if err != nil {
			panic(err)
		}
		tenantId := config.Auth.TenantId
		tokenId := config.Token.Id

		servers := &[]endpoint.Server{}
		err = endpoint.Get(tenantId, tokenId, servers)
		if err != nil {
			panic(err)
		}

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

		err = endpoint.Delete(tenantId, tokenId, id)
		if err != nil {
			panic(err)
		}
		fmt.Println("delete succesful.")
	},
}
