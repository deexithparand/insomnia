package insomnia

import (
	"insomnia/internal"
	"insomnia/utils"
)

// read config from file
func ReadConfig(configFilePath string) {
	utils.YMLParser(configFilePath)
}

// Setup DB
func DB() {
	internal.TestPingDB()
}

func init() {
	// fmt.Println("First init() function in main package")
	DB()
}
