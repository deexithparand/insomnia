package internal

import (
	"database/sql"
	"fmt"
	"insomnia/utils"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.215.2"
	port     = 5432
	user     = "postgres"
	password = "dev"
	dbname   = "insomnia_dev"
)

var DB *sql.DB

func sqlFromFile(file string) string {
	base := "internal"
	path := filepath.Join(base, file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
	return string(data)
}

func Migrate() {
	// DB Migration Code
	err := DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("If we ping we can also migrate")

	// Table creations
	log.Println("Starting Migrations ... ")

	log.Println("Table Migrations ... ")

	result, err := DB.Exec(sqlFromFile("create-table.sql"))
	if err != nil {
		panic(err)
	}

	log.Println("Tables Created : ", result)
}

func Seed(config utils.Config) {
	// Seed Initial Config to DB
}

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database context created & connected successfully")
}

func Close() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Closed connection to DB successfully")
}
