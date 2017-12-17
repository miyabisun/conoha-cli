package cmd

import (
	"bytes"
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
	RootCmd.AddCommand(UpCmd)
}

var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "up in ConoHa API.",
	Long:  "up in ConoHa API(required logged in)",
	Run: func(cmd *cobra.Command, args []string) {
		spec := &Spec{}
		_, err := toml.DecodeFile("spec.toml", spec)
		if err != nil {
			panic(err)
		}
		config, _ := conf.Read()
		tokenId, _ := util.TokenId()

		// TODO:ステータス取得コードはリファクタリング
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

		var name string
		for _, it := range vmlist.Servers {
			if it.Metadata.Instance_name_tag == spec.Name {
				name = it.Metadata.Instance_name_tag
				break
			}
		}
		if name != "" {
			fmt.Printf("%s server is alive.\n", name)
			return
		}

		var image_id string
		images := &[]util.Item{}
		util.Info("images", images)
		for _, it := range *images {
			if it.Name == spec.Image {
				image_id = it.Id
				break
			}
		}
		if image_id == "" {
			fmt.Printf("image %s is not found.\n", spec.Image)
			return
		}

		var flavor_id string
		flavors := &[]util.Item{}
		util.Info("flavors", flavors)
		for _, it := range *flavors {
			if it.Name == spec.Flavor {
				flavor_id = it.Id
				break
			}
		}
		if flavor_id == "" {
			fmt.Printf("flavor %s is not found.\n", spec.Flavor)
			return
		}

		reqBody := fmt.Sprintf(`{"server": {"imageRef": "%s", "flavorRef": "%s", "metadata": {"instance_name_tag": "%s"}, "key_name": "%s"} }`, image_id, flavor_id, spec.Name, spec.SSHKey)
		buffer := bytes.NewBufferString(reqBody)
		url = fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/servers", config.Auth.TenantId)
		client = &http.Client{}
		request, _ = http.NewRequest("POST", url, buffer)
		request.Header.Add("Accept", "application/json")
		request.Header.Add("X-Auth-Token", tokenId)
		resp2, _ := client.Do(request)
		defer resp2.Body.Close()
		body, _ = ioutil.ReadAll(resp2.Body)
	},
}
