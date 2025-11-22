package insomnia

import (
	"insomnia/internal"
	"insomnia/internal/core"
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

func Start() {
	// Start Monitoring, Wait & Trigger
	core.Monitor()
}
