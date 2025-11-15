package insomnia

import "insomnia/utils"

func ReadConfig(configFilePath string) {
	utils.YMLParser(configFilePath)
}
