package database

import "log"

// seeder to workspaces (input must : as an object)
func (db *Database) SeedWorkspaces() {
	workspaces := []Workspace{
		{ID: "e2ewdw", Name: "Workspace1"},
		{ID: "32dwdw", Name: "Workspace2"},
		{ID: "32dfefdwwd", Name: "Workspace3"},
	}

	for _, workspace := range workspaces {
		result := db.Conn.Create(&workspace)
		if result.Error != nil {
			log.Printf("error seeding the database : %v", result.Error)
		}
	}

	log.Printf("Successfully seeded workspace database")

}
