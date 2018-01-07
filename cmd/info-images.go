package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/images"
	"github.com/spf13/cobra"
)

var InfoImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "get images ConoHa API.",
	Long:  "get images ConoHa API (require logged in).",
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
		tokenId := config.Token.Id

		images := &[]endpoint.Image{}
		err = endpoint.Get(tokenId, images)
		if err != nil {
			panic(err)
		}

		for _, item := range *images {
			fmt.Println(item.Name)
		}
	},
}

func init() {
	InfoCmd.AddCommand(InfoImagesCmd)
}
