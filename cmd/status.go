package cmd

import (
	"fmt"
	"os"

	"github.com/miyabisun/conoha-cli/config/conoha"
	"github.com/miyabisun/conoha-cli/config/spec"
	"github.com/miyabisun/conoha-cli/config/status"
	endpoint "github.com/miyabisun/conoha-cli/endpoints/servers"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var isAll bool

func init() {
	RootCmd.AddCommand(StatusCmd)
	StatusCmd.Flags().BoolVarP(&isAll, "all", "a", false, "Show All Status")
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status in ConoHa API.",
	Long:  "status in ConoHa API(required logged in)",
	Run: func(cmd *cobra.Command, args []string) {
		try := util.Try
		try(conoha.Refresh())

		config := &conoha.Config{}
		try(conoha.Read(config))
		tenantId := config.Auth.TenantId
		tokenId := config.Token.Id

		confSpec := &spec.Config{}
		try(spec.Read(confSpec))
		name := confSpec.Name

		server := &endpoint.Server{}
		try(endpoint.Show(tenantId, tokenId, name, server))

		state := server.Status
		if state == "" {
			state = "NONE"
		} else {
			try(status.Save(server))
		}

		data := [][]string{
			[]string{server.Metadata.Instance_name_tag, state},
		}

		if isAll {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "State"})
			table.SetBorder(false)
			table.AppendBulk(data)
			table.Render()
		} else {
			fmt.Println(state)
		}
	},
}
