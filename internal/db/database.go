package db

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

// *** dependency functions ***
func intervalSQLQuery(interval int, unitShorthand string) string {

	unitMap := map[string]string{
		"s": "SECONDS",
		"m": "MINUTES",
		"h": "HOURS",
	}

	return fmt.Sprintf("%d %s", interval, unitMap[unitShorthand])
}

func sqlFromFile(baseFolder, file string) string {
	base := "internal/db/sql"
	path := filepath.Join(base, baseFolder, file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
	return string(data)
}

// *** core functions ***
func Migrate() {

	// Table creations
	log.Println("Starting Migrations ... ")

	result, err := DB.Exec(sqlFromFile("migrate", "migrate.sql"))
	if err != nil {
		panic(err)
	}

	log.Println("Migrations Complete ... ", result)
}

func Seed(config utils.Config) {

	// seeding targetgroup table
	for _, targetgroup := range config.Insomnia.TargetGroups {
		_, err := DB.Query(sqlFromFile("seed", "targetgroup.sql"), targetgroup.Label)
		if err != nil {
			panic(err)
		}
	}
	log.Println("Seeded targetgroups from config file")

	// seeding target table
	for _, targetgroup := range config.Insomnia.TargetGroups {
		// load target group wise
		label := targetgroup.Label
		for _, target := range targetgroup.Targets {
			_, err := DB.Query(sqlFromFile("seed", "target.sql"), label, target.Endpoint.Url, intervalSQLQuery(target.Endpoint.Interval, target.Endpoint.Unit))
			if err != nil {
				panic(err)
			}
		}
		log.Println("Loaded targets of the targetgroup : ", label)
	}
	log.Println("Seeded targets from config file")
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
