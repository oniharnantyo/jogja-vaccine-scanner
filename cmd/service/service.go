package service

import (
	"github.com/oniharnantyo/jogja-vaccine-scanner/config"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/bootstrap"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

var command = &cobra.Command{
	Use:     "service",
	Aliases: []string{"svc"},
	Short:   "Run service",
	Run: func(c *cobra.Command, args []string) {
		conf, err := config.Load(configFile)
		if err != nil {
			panic(err)
		}
		bootstrap.Run(conf)
	},
}

func init() {
	command.Flags().StringVar(&configFile, "config", "./config.yaml", "Set config file path")
}

func GetCommand() *cobra.Command {
	return command
}
