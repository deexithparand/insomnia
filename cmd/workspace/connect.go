package workspace

import (
	"fmt"
	"insomnia/format"
	"insomnia/helper"
	"insomnia/state"

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

type workspaceConnectInput struct {
	connectedWorkspaceName string
}

func generateWorkspaceNamesTable() {
	var workspaceNamesTable [][]string

	// fetch already present workspaces
	workspaceData, err := state.DB.GetAllWorkspaces()
	if err != nil {
		fmt.Printf("error retrieving workspaces : %v", err)
		return
	}

	// Name of workspaces in seperate slice
	for _, singleWorkspaceData := range workspaceData {
		// getting names from index 1
		var workspaceNames []string
		workspaceNames = append(workspaceNames, singleWorkspaceData[1])
		workspaceNamesTable = append(workspaceNamesTable, workspaceNames)
	}

	// list the workspace names as tables
	tableHeaders := []string{
		"Workspaces",
	}

	format.TableFormat(tableHeaders, workspaceNamesTable)
}

func connectToWorkspace() {
	var input workspaceConnectInput
	var err error

	// list all the available workspace
	fmt.Println("Available workspaces : ")
	generateWorkspaceNamesTable()

	// get the user input
	input.connectedWorkspaceName, err = helper.GetCMDInput("Name the workspace to which you would like to connect : ")
	if err != nil {
		fmt.Printf("error getting the CMD input : %v\n", err)
		return
	}

	// update the current connected workspace
	state.SetCurrentWorkspace(input.connectedWorkspaceName)

	// display the updated workspace from state
	fmt.Printf("Successfully connected to workspace : %s", state.GetCurrentWorkspace())

}

func init() {
	// rootCmd.AddCommand()
}
