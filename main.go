package main

import "insomnia/insomnia"

func main() {

	// config values
	configFilePath := "./config.test.yml"

	// read config file
	insomnia.ReadConfig(configFilePath)
}
