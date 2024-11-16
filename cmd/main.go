package main

import (
	"insomina/config"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// load the config from env
	var (
		c           config.Config
		environment string
	)

	environment = "local"

	err := c.LoadConfig(environment)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	

}
