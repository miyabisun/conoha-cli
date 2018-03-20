package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	"github.com/miyabisun/conoha-cli/endpoints/flavors"
	"github.com/miyabisun/conoha-cli/endpoints/images"
	"github.com/miyabisun/conoha-cli/endpoints/servers"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(UpCmd)
}

var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "VPSインスタンスの起動",
	Long:  "up in ConoHa API(required logged in)",
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

		server := &servers.Server{}
		try(servers.Show(tenantId, tokenId, name, server))

		serverName := server.Metadata.Instance_name_tag
		if serverName != "" {
			fmt.Printf("%s server is alive.\n", name)
			return
		}

		image := &images.Image{}
		try(images.Show(tokenId, confSpec.Image, image))
		if image.Id == "" {
			fmt.Printf("image %s is not found.\n", confSpec.Image)
			return
		}

		flavor := &flavors.Flavor{}
		try(flavors.Show(tenantId, tokenId, confSpec.Flavor, flavor))
		if flavor.Id == "" {
			fmt.Printf("flavor %s is not found.\n", confSpec.Flavor)
			return
		}

		try(servers.Post(tenantId, tokenId, name, image.Id, flavor.Id, confSpec.SSHKey))
		fmt.Println("up successful.")
	},
}
