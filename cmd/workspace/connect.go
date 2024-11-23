package workspace

import (
	"fmt"
	"insomnia/format"
	"insomnia/helper"
	"insomnia/state"

	"github.com/manifoldco/promptui"
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
	requireWorkspace       string
}

func generateWorkspaceNamesTable() {
	var workspaceNames []string
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
		workspaceNames = append(workspaceNames, singleWorkspaceData[1])
	}
	workspaceNamesTable = append(workspaceNamesTable, workspaceNames)

	// list the workspace names as tables
	tableHeaders := []string{
		"Workspaces",
	}

	format.TableFormat(tableHeaders, workspaceNamesTable)
}

func connectToWorkspace() {
	var input workspaceConnectInput
	var err error

	// get the currently connected workspace
	if len(input.connectedWorkspaceName) == 0 {
		fmt.Println("Not connected to any workspace yet")
	} else {
		fmt.Println("Currently connected workspace : ", input.connectedWorkspaceName)
	}

	// Connect to a workspace ?
	prompt := promptui.Select{
		Label: "Would you like to connect to new workspace ?",
		Items: []string{"Yes", "No"},
	}

	_, input.requireWorkspace, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if input.requireWorkspace == "Yes" {

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
		fmt.Printf("Successfully updated the connected workspace : %s", state.GetCurrentWorkspace())
	}

}

func init() {
	// rootCmd.AddCommand()
}
