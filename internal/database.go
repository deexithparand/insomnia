package internal

import (
	"database/sql"
	"fmt"
	"insomnia/utils"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.215.2"
	port     = 5432
	user     = "postgres"
	password = "dev"
	dbname   = "insomnia_dev"
)

func TestPingDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database connected successfully")
}

func Migrate() {
	// DB Migration Code
}

func Seed(config utils.Config) {
	// Seed Initial Config to DB
}
