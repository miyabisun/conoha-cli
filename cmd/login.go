package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/conf"
	"github.com/miyabisun/conoha-cli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(LoginCmd)
	LoginCmd.Flags().StringP("username", "u", "", "ユーザー名 (gncu00000000)")
	LoginCmd.Flags().StringP("password", "p", "", "パスワード (9文字以上)")
	LoginCmd.Flags().StringP("tenantid", "t", "", "テナントID (半角英数32文字)")
	viper.BindPFlag("auth.username", LoginCmd.Flags().Lookup("username"))
	viper.BindPFlag("auth.password", LoginCmd.Flags().Lookup("password"))
	viper.BindPFlag("auth.tenant_id", LoginCmd.Flags().Lookup("tenantid"))
}

func findAuth() conf.ConfAuth {
	auth := conf.ConfAuth{
		User:     viper.GetString("auth.username"),
		Pass:     viper.GetString("auth.password"),
		TenantId: viper.GetString("auth.tenant_id"),
	}
	fmt.Printf("auth: %T, %s\n", auth, auth)
	if auth.User == "" {
		fmt.Print("username: ")
		fmt.Scan(&auth.User)
	}
	if auth.Pass == "" {
		fmt.Print("password: ")
		fmt.Scan(&auth.Pass)
	}
	if auth.TenantId == "" {
		fmt.Print("tenant_id: ")
		fmt.Scan(&auth.TenantId)
	}
	return auth
}

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to ConoHa API.",
	Long:  "login to ConoHa API.",
	Run: func(cmd *cobra.Command, args []string) {
		auth := findAuth()
		body, statusCode, err := util.LoginFrom(&auth)
		if err != nil {
			panic(err)
		}
		fmt.Printf("statusCode: %s\n", statusCode)
		fmt.Println(string(body))
		if statusCode != 200 {
			return
		}

		token, err := util.ToToken(body)
		if err != nil {
			panic(err)
		}

		config, err := conf.Read()
		if err != nil {
			panic(err)
		}
		config.Auth = auth
		config.Token = token

		err = conf.Write(&config)
		if err != nil {
			panic(err)
		}
		fmt.Println("login successful.")
	},
}
