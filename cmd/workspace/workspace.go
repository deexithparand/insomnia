package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var WorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "workspace is the parent command, used to point all workspace related commands and subcommands",
	Long:  `workspace is the parent command, used to point all workspace related commands and subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No subcommand provided. Use 'insomnia workspace --help' for more information.")
	},
}

func init() {
	WorkspaceCmd.AddCommand(listCmd)
	WorkspaceCmd.AddCommand(connectCmd)
	WorkspaceCmd.AddCommand(createCmd)
}
