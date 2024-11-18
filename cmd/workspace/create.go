package workspace

import (
	"bufio"
	"fmt"
	"insomnia/state"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new workspace",
	Long:  `create a new workspace, which basically helps to group a set of endpoints or app urls to be monitored`,
	Run: func(cmd *cobra.Command, args []string) {
		createWorkspace()
	},
}

func createWorkspace() {
	var workspaces []string

	// fetch already present workspaces
	workspaces, err := state.DB.GetAllWorkspaces()
	if err != nil {
		fmt.Printf("error retrieving workspaces : %v", err)
		return
	}

	// prompt for a workspace name
	reader := bufio.NewReader(os.Stdin)
	workspacename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading the new workspace name input : %v", err)
		return
	}

	// check if the workspace is there

}

func init() {
	// rootCmd.AddCommand()
}
