package database

import (
	"insomnia/helper"
	"log"
)

// seeder to workspaces (input must : as an object)
func (db *Database) SeedWorkspaces() {
	workspaces := []Workspace{
		{ID: helper.GenerateUUID(), Name: "testworkspace"},
	}

	for _, workspace := range workspaces {
		result := db.Conn.Create(&workspace)
		if result.Error != nil {
			log.Printf("error seeding the database : %v", result.Error)
		}
	}

	log.Printf("Successfully seeded workspace database")

}
