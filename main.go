package main

import (
	"insomnia/cmd"
	"insomnia/config"
	"insomnia/database"
	"insomnia/state"
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
		panic(err)
	}

	// configure db connections
	database.ConfigureDSN(cfg)
	err = database.Connect()
	if err != nil {
		panic(err)
	}

	// migrate db
	database.Migrate()
	database.SeedWorkspaces()

	// cli
	state.SetDatabase(database)
	cmd.Execute()

}
