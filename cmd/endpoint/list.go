package endpoint

import (
	"fmt"
	"insomnia/state"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the endpoints (from a connected project)",
	Long:  `lists all the endpoints (from a connected project)`,
	Run: func(cmd *cobra.Command, args []string) {
		listEndpoints()
	},
}

func listEndpoints() {
	fmt.Println("Here is the list of endpoints from the workspace : ", state.GetCurrentWorkspace())
}

func init() {
	// rootCmd.AddCommand()
}
