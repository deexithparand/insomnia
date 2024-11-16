package database

import (
	"fmt"
	"insomnia/config"

	"gorm.io/gorm"
)

type Database struct {
	DSN  string
	Conn *gorm.DB
}

// setup db config - dsn (data source name)
func (db *Database) ConfigureDSN(dbconf config.Config) {
	db.DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		dbconf.DB_URL,
		dbconf.DB_USER,
		dbconf.DB_PASSWORD,
		dbconf.DB_NAME,
		dbconf.DB_PORT,
		dbconf.DB_SSLMODE,
	)
}
