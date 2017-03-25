package main

type glide struct {
	Name    string       `yaml:"package"`
	License string       `yaml:"license"`
	Imports dependencies `yaml:"import"`
}

type dependencies []dependency

type dependency struct {
	Package string `yaml:"package"`
	Version string `yaml:"version"`
}

type manifest struct {
	Dependencies map[string]manifestValue `json:"dependencies,omitempty"`
}

type manifestValue struct {
	Branch  string `json:"branch,omitempty"`
	Version string `json:"version,omitempty"`
}

const glideFile = "glide.yaml"
const manifestFile = "mainfest.json"

func main() {

}
