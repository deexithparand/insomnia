package main

import (
	"insomnia/cmd"
	"insomnia/config"
	"insomnia/database"
	"insomnia/state"
	"log"
)

func main() {
	// load the config from env
	var (
		cfg      config.Config
		database database.Database
	)

	// environment constant
	const env string = "local"

	// load config from env
	err := cfg.LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// configure db connections
	database.ConfigureDSN(cfg)
	err = database.Connect()
	if err != nil {
		log.Fatalf("error connecting to database : %v", err)
	}

	// migrate db
	database.Migrate()
	database.SeedWorkspaces()

	// cli
	state.SetDatabase(database)
	cmd.Execute()

}
