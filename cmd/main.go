package main

import (
	"insomnia/config"
	db "insomnia/db/config"

	"log"
)

func main() {
	// load the config from env
	var (
		cfg config.Config
		db  db.Database
	)

	// environment constant
	const env string = "local"

	// load config from env
	err := cfg.LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// db migrations

	// configure db connections
	db.ConfigureDSN(cfg)
	err = db.Connect()
	if err != nil {
		log.Fatalf("error connecting to database : %v", err)
	}

}
