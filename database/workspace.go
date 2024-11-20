package database

import (
	"fmt"
)

func (db *Database) GetAllWorkspaces() ([][]string, error) {

	var workspaces []Workspace

	// get all workspaces
	result := db.Conn.Find(&workspaces)
	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving records from workspace table : %v", result.Error)
	}

	var workspaceData [][]string

	for _, workspace := range workspaces {
		// fresh records
		var records []string

		// Get ID
		records = append(records, workspace.ID)

		// Get Name
		records = append(records, workspace.Name)

		// Get Description
		records = append(records, workspace.Description)

		workspaceData = append(workspaceData, records)
	}

	return workspaceData, nil
}
