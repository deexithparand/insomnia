package insomnia

import (
	"insomnia/internal/core"
	"insomnia/internal/db"
	"insomnia/utils"
)

// read config from file
func Config(configFilePath string) utils.Config {
	// Parse and Return Config
	return utils.YMLParser(configFilePath)
}

// Setup DB
func DB(config utils.Config) {
	// Load Data To DB

	db.Connect()
	db.Migrate()
	db.Seed(config)
}

func AppStart() {
	// Start Monitoring, Wait & Trigger
	core.Monitor()
}
