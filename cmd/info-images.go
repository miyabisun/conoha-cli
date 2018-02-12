package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/images"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
)

var InfoImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "get images ConoHa API.",
	Long:  "get images ConoHa API (require logged in).",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try
		try(conoha.Refresh())

		config := &conoha.Config{}
		try(conoha.Read(config))
		tokenId := config.Token.Id

		images := &[]endpoint.Image{}
		try(endpoint.Get(tokenId, images))

		for _, item := range *images {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoImagesCmd)
}
