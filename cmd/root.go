package root

import (
	"github.com/oniharnantyo/jogja-vaccine-scanner/cmd/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "template",
}

func init() {
	rootCmd.AddCommand(service.GetCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
