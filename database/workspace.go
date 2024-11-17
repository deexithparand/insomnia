package database

import (
	"fmt"
)

func (db *Database) GetAllWorkspaces() ([]string, error) {

	var records []string
	var workspaces []Workspace

	// get all workspaces
	result := db.Conn.Find(&workspaces)
	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving records from workspace table : %v", result.Error)
	}

	for _, workspace := range workspaces {
		records = append(records, workspace.Name)
	}

	return records, nil
}
