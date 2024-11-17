package state

import "insomnia/database"

var DB database.Database

func SetDatabase(database database.Database) {
	DB = database
}
