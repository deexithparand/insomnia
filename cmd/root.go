package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "insomnia",
	Short: "Insomnia - A CLI tool for monitoring uptime and endpoints",
	Long: `Insomnia is a lightweight CLI tool designed for uptime monitoring.
It helps you manage workspaces and monitor endpoints effortlessly.
Easily track, create, and manage server and service availability
with an intuitive command-line interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
