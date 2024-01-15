package cmd

import (
	"backend-api/app/config"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	EnvFilePath string
	rootCmd     = &cobra.Command{
		Use:   "cobra-cli",
		Short: "backend api",
	}
)
var (
	rootConfig *config.Root
)

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&EnvFilePath, "env", "e", ".env", ".env file to read from")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Cannot Run CLI. err > ", err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(func() {
		configReader()
	})
}
func configReader() {
	log.Infof("Initialize ENV")
	rootConfig = config.Load(EnvFilePath)
}
