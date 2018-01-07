package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/servers"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(StatusCmd)
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status in ConoHa API.",
	Long:  "status in ConoHa API(required logged in)",
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

		confSpec := &spec.Config{}
		err = spec.Read(confSpec)
		if err != nil {
			panic(err)
		}
		name := confSpec.Name

		server := &endpoint.Server{}
		err = endpoint.Show(tenantId, tokenId, name, server)
		if err != nil {
			panic(err)
		}

		status := server.Status
		if status == "" {
			status = "NONE"
		}
		fmt.Println(status)
	},
}
