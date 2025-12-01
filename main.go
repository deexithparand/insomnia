package main

import (
	"fmt"
	"insomnia/internal/core"
	"insomnia/internal/db"
	"insomnia/internal/utils"
	"os"
	"os/signal"
	"syscall"
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

func main() {

	// Config Scripts
	configFilePath := "./config.test.yml"
	config := Config(configFilePath)

	// DB Scripts
	DB(config)
	defer db.Close()

	AppStart()

	// keeps the application running - until we terminate
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("App Started - Runs until interrupt")
	<-done
}
