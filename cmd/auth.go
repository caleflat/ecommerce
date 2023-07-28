package cmd

import (
	"github.com/caleflat/ecommerce/internal/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use: "auth",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Execute()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
