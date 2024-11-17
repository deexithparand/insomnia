package endpoint

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var EndpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "endpoint is the parent command, used to point all endpoint related commands and subcommands",
	Long:  `endpoint is the parent command, used to point all endpoint related commands and subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No subcommand provided. Use 'insomnia endpoint --help' for more information.")
	},
}

func init() {
	EndpointCmd.AddCommand(listCmd)
}
