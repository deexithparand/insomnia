package database

import "log"

func (db *Database) Migrate() {
	db.Conn.AutoMigrate(&Workspace{})
	db.Conn.AutoMigrate(&Endpoint{})

	log.Println("Tables migrated successfully")
}
