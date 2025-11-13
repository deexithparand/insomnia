package internal

import "insomnia/utils"

type Endpoint struct {
	URL      string
	Interval string
}

type Target struct {
	Endpoints []Endpoint
}

type TargetGroup struct {
	Label   string
	Targets []Target
}

type Insomnia struct {
	TG TargetGroup
}

func ReadConfig() {
	// I/P the config.test.yml
	// Make the objects accessible and print here
	utils.ParserTest()
}
