package utils

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
