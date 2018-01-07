package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	"github.com/miyabisun/conoha-cli/endpoints/flavors"
	"github.com/miyabisun/conoha-cli/endpoints/images"
	"github.com/miyabisun/conoha-cli/endpoints/servers"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(UpCmd)
}

var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "up in ConoHa API.",
	Long:  "up in ConoHa API(required logged in)",
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

		server := &servers.Server{}
		err = servers.Show(tenantId, tokenId, name, server)
		if err != nil {
			panic(err)
		}

		serverName := server.Metadata.Instance_name_tag
		if serverName != "" {
			fmt.Printf("%s server is alive.\n", name)
			return
		}

		image := &images.Image{}
		err = images.Show(tokenId, confSpec.Image, image)
		if err != nil {
			panic(err)
		}
		if image.Id == "" {
			fmt.Printf("image %s is not found.\n", confSpec.Image)
			return
		}

		flavor := &flavors.Flavor{}
		err = flavors.Show(tenantId, tokenId, confSpec.Flavor, flavor)
		if err != nil {
			panic(err)
		}
		if flavor.Id == "" {
			fmt.Printf("flavor %s is not found.\n", confSpec.Flavor)
			return
		}

		err = servers.Post(tenantId, tokenId, name, image.Id, flavor.Id, confSpec.SSHKey)
		if err != nil {
			panic(err)
		}
		fmt.Println("up successful.")
	},
}
