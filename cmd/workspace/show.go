package workspace

import (
	"fmt"
	"insomnia/state"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the name of the currently connected workspace",
	Long:  `create a new workspace, which basically helps to group a set of endpoints or app urls to be monitored`,
	Run: func(cmd *cobra.Command, args []string) {
		showConnectedDatabase()
	},
}

func showConnectedDatabase() {
	// set the workspace name
	var connectedWorkspaceName string = state.GetCurrentWorkspace()

	// get the currently connected workspace
	if len(connectedWorkspaceName) == 0 {
		fmt.Println("Not connected to any workspace yet")
	} else {
		fmt.Println("Currently connected workspace : ", connectedWorkspaceName)
	}
}

func init() {
	// rootCmd.AddCommand()
}
