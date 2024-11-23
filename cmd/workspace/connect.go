package workspace

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the workspace specified",
	Long: `helps to connect to an already present workspace
	so that it will be easy to monitor the applications under each workspace granularly`,
	Run: func(cmd *cobra.Command, args []string) {
		connectToWorkspace()
	},
}

func connectToWorkspace() {
	// list all the available workspace

	// get the user input

	// check if the workspace is already created

	// if so update the currentworkspace to that one
}

func init() {
	// rootCmd.AddCommand()
}
