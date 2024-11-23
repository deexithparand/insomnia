package workspace

import (
	"errors"
	"fmt"
	"insomnia/helper"
	"insomnia/state"

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

type Workspace struct {
	ID          string
	Name        string
	Description string
}

type Input struct {
	workspaceName        string
	workspaceDescription string
}

func createWorkspace() {
	var workspaceNames []string
	var workspaceData [][]string
	var err error

	var input Input

	// fetch already present workspaces
	workspaceData, err = state.DB.GetAllWorkspaces()
	if err != nil {
		fmt.Printf("error retrieving workspaces : %v", err)
		return
	}

	// Name of workspaces in seperate slice
	for _, workspace := range workspaceData {
		// getting names from index 1
		workspaceNames = append(workspaceNames, workspace[1])
	}

	for {

		// prompt for a workspace name
		input.workspaceName, err = helper.GetCMDInput("Enter workspace name : ")
		if err != nil {
			fmt.Printf("error getting the input from cmd : %v\n", err)
			return
		}

		// impose CMD ruleset
		followsRuleset := helper.CMDInputRuleSet(input.workspaceName)
		if !followsRuleset {
			followsRulesetError := errors.New("only lowercase letters, number and hyphens allowed for this input ")
			fmt.Printf("error : %v\n", followsRulesetError)
			continue
		}

		// input workspace description
		input.workspaceDescription, err = helper.GetCMDInput("Enter workspace description : ")
		if err != nil {
			fmt.Printf("error getting the input from cmd : %v\n", err)
			continue
		}

		// check if the workspace is there
		var isPresent bool = false
		for _, workspaceName := range workspaceNames {
			if workspaceName == input.workspaceName {
				isPresent = true
				break
			}
		}

		// is already present throw error and reinstantiate the question
		if isPresent {
			workspaceNameTakenError := errors.New("workspace name is already taken, try a new one")
			fmt.Printf("error : %v\n", workspaceNameTakenError)
			continue
		} else {
			break
		}

	}

	// insert a new workspace with the name

	var newWorkspace Workspace = Workspace{
		ID:          helper.GenerateUUID(),
		Name:        input.workspaceName,
		Description: input.workspaceDescription,
	}

	result := state.DB.Conn.Create(&newWorkspace)
	if result.Error != nil {
		fmt.Printf("error creating new workspace in the db : %v", result.Error)
		return
	}

	fmt.Printf("Successfully created a new workspace named : %s\n", input.workspaceName)
}

func init() {
	// rootCmd.AddCommand()
}
