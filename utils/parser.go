package utils

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type Endpoint struct {
	Url      string `yaml:"url"`
	Interval string `yaml:"interval"`
}

type Target struct {
	Endpoint Endpoint `yaml:"endpoint"`
}

type TargetGroup struct {
	Label   string   `yaml:"label"`
	Targets []Target `yaml:"targets"`
}

type Insomnia struct {
	TargetGroups []TargetGroup `yaml:"target-group"`
}

type Config struct {
	Insomnia Insomnia `yaml:"insomnia"`
}

// DEV Usage - For Logging Parsed Output
// func YMLParserLogger(config Config) {
// 	// looping through all target groups
// 	for index, tg := range config.Insomnia.TargetGroups {

// 		log.Printf("Running Target Group (%d) : %s ...\n", index, strings.ToUpper(tg.Label))

// 		// looping through - endpoints and their intervals in the target group
// 		for _, target := range tg.Targets {
// 			log.Println("Endpoint : ", target.Endpoint.Url, " - ", target.Endpoint.Interval)
// 		}

// 		fmt.Println() // line space after every target group log
// 	}

// }

func YMLParser(path string) Config {

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	rawConfigContent := string(fileBytes)

	var config Config
	err = yaml.Unmarshal([]byte(rawConfigContent), &config)
	if err != nil {
		panic(err)
	}

	log.Println("Config file parsed successfully")

	return config
}
