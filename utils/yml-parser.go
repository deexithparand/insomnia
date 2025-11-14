package utils

import (
	"fmt"

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

func ParserTest() {
	rawConfigContent := `
# This is a config file for the insomnia - endpoints & their properties
insomnia:
  target-group:
      - label:  "sample"
        targets: 
          - endpoint:
              url: "https://www.google.com"    
              interval: 5m # only whole values ( decimals not allowed ) - in s, m, h, d ( default - 5m )
          - endpoint:
              url: "https://www.whatismyip.com"
              interval: 5s
        # global-setting: ( planning to add this later )
`
	var config Config
	err := yaml.Unmarshal([]byte(rawConfigContent), &config)
	if err != nil {
		panic(err)
	}

	// looping through all target groups
	for _, tg := range config.Insomnia.TargetGroups {

		// currently only one target group
		fmt.Println("Target Group : ", tg.Label)
	}

}
