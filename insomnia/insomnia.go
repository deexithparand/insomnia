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
	// Test Connection
	internal.TestPingDB()

	// Load Data To DB
	internal.Migrate()
	internal.Seed(config)
}

func Monitor() {
	// currently kept empty until migration
}

func Start() {

	// config values
	configFilePath := "./config.test.yml"

	// Config Scripts
	config := Config(configFilePath)

	// DB Scripts
	DB(config)

	// Start Monitoring, Wait & Trigger

}
