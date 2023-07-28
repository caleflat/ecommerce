package cmd

import (
	"github.com/caleflat/ecommerce/internal/serve"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		serve.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
