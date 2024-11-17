package workspace

import (
	"fmt"
	"insomnia/state"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the workspaces",
	Long: `lists all the workspaces created by the user. 
	Using the name of the workspace we can track down the endpoints of the apps to be monitored under each project`,
	Run: func(cmd *cobra.Command, args []string) {
		listWorkspaces()
	},
}

func listWorkspaces() {
	workspaces, err := state.DB.GetAllWorkspaces()
	if err != nil {
		fmt.Printf("error retrieving the workspaces from the database : %v", err)
		return
	}

	fmt.Println("Workspaces       |")
	fmt.Println("------------------")

	for _, workspace := range workspaces {
		fmt.Println(workspace)
	}
}

func init() {
	// rootCmd.AddCommand()
}
