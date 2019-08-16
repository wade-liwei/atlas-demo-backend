package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wade-liwei/atlas-demo-backend/controller"
	//homedir "github.com/mitchellh/go-homedir"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "atlas demo backend",

	Run: func(cmd *cobra.Command, args []string) {
		//config
		v := viper.GetViper()
		log_level_str := v.GetString("log_level")
		if len(log_level_str) == 0 {
			log_level_str = "cmd:debug,app:debug,*:error"
		}

		controller.Start()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
