package workspace

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var WorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "workspace is the parent command, used to point all workspace related commands and subcommands",
	Long:  `workspace is the parent command, used to point all workspace related commands and subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("No subcommand provided. Use 'insomnia workspace --help' for more information.")
		cmd.Help()
	},
}

func init() {
	WorkspaceCmd.AddCommand(connectCmd)
	WorkspaceCmd.AddCommand(createCmd)
	WorkspaceCmd.AddCommand(listCmd)
	WorkspaceCmd.AddCommand(showCmd)
}
