package database

import (
	"errors"
	"insomnia/helper"
)

// seeder to workspaces (input must : as an object)
func (db *Database) SeedWorkspaces() error {
	workspaces := []Workspace{
		{ID: helper.GenerateUUID(), Name: "testworkspace"},
	}

	for _, workspace := range workspaces {
		result := db.Conn.Create(&workspace)
		if result.Error != nil {
			return errors.New("error seeding the database : " + result.Error.Error())
		}
	}

	return nil
}
