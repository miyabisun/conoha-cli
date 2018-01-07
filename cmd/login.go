package cmd

import (
	"fmt"

	"github.com/miyabisun/conoha-cli/config/conoha"
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

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to ConoHa API.",
	Long:  "login to ConoHa API.",
	Run: func(cmd *cobra.Command, args []string) {
		config := &conoha.Config{}
		err := conoha.Read(config)
		if err != nil {
			panic(err)
		}
		config.Auth = *findAuth()
		fmt.Printf("auth: %T, %s\n", config.Auth, config.Auth)

		err = conoha.Login(&config.Auth, &config.Token)
		if err != nil {
			panic(err)
		}

		err = conoha.Write(config)
		if err != nil {
			panic(err)
		}
		fmt.Println("login successful.")
	},
}

func findAuth() *conoha.ConfigAuth {
	auth := conoha.ConfigAuth{
		User:     viper.GetString("auth.username"),
		Pass:     viper.GetString("auth.password"),
		TenantId: viper.GetString("auth.tenant_id"),
	}
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
	return &auth
}
