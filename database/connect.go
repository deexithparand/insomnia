package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// open connection to db (postgres just for now)
func (db *Database) Connect() error {
	conn, err := gorm.Open(postgres.Open(db.DSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening connection to database : %v", err.Error())
	}

	db.Conn = conn
	return nil
}
