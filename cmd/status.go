package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/miyabisun/conoha-cli/conf"
	"github.com/miyabisun/conoha-cli/util"
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
		spec := &Spec{}
		_, err := toml.DecodeFile("spec.toml", spec)
		if err != nil {
			panic(err)
		}
		name := spec.Name
		config, _ := conf.Read()
		tokenId, _ := util.TokenId()

		url := fmt.Sprintf(`https://compute.tyo1.conoha.io/v2/%s/servers/detail`, config.Auth.TenantId)
		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Add("Accept", "application/json")
		request.Header.Add("X-Auth-Token", tokenId)
		resp, _ := client.Do(request)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		vmlist := &Vmlist{}
		json.Unmarshal(body, vmlist)

		var status string
		for _, server := range vmlist.Servers {
			if server.Metadata.Instance_name_tag == name {
				status = server.Status
			}
		}
		if status == "" {
			status = "NONE"
		}
		fmt.Println(status)
	},
}

type Vmlist struct {
	Servers []Server
}
type Server struct {
	Id       string
	Status   string
	Image    struct{ Id string }
	Flavor   struct{ Id string }
	Metadata struct{ Instance_name_tag string }
}
type Spec struct {
	Name   string
	Image  string
	Flavor string
	SSHKey string
}
