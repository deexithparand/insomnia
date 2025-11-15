package insomnia

import (
	"insomnia/internal"
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
	internal.Connect()
	internal.Migrate()
	internal.Seed(config)
}

func Monitor() {
	// currently kept empty until migration

	// Table Creation

}

func Start() {
	// Start Monitoring, Wait & Trigger

}
