package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "conoha",
	Short: "This tool is conoha cli.",
	Long:  "This tool is a conoha cli.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(versionCmd)
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigName("conoha")
	viper.AddConfigPath("$HOME/.config")
	viper.ReadInConfig()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of conoha-cli",
	Long:  `All software has versions. This is conoha-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("conoha-cli v0.0.1")
	},
}
