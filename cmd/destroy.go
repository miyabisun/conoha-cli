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
	RootCmd.AddCommand(DestroyCmd)
}

var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy in ConoHa API.",
	Long:  "destroy in ConoHa API(required logged in)",
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
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Add("Accept", "application/json")
		request.Header.Add("X-Auth-Token", tokenId)
		resp, _ := client.Do(request)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		vmlist := &Vmlist{}
		json.Unmarshal(body, vmlist)

		var id string
		for _, it := range vmlist.Servers {
			if it.Metadata.Instance_name_tag == name {
				id = it.Id
			}
		}
		if id == "" {
			fmt.Println("not found server.")
			return
		}

		url = fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/servers/%s", config.Auth.TenantId, id)
		client = &http.Client{}
		request, _ = http.NewRequest("DELETE", url, nil)
		request.Header.Add("Accept", "application/json")
		request.Header.Add("X-Auth-Token", tokenId)
		resp2, _ := client.Do(request)
		defer resp2.Body.Close()
		body, _ = ioutil.ReadAll(resp2.Body)
	},
}
