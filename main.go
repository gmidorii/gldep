package main

import (
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type glide struct {
	Name    string       `yaml:"package"`
	Imports dependencies `yaml:"import"`
}

func (g glide) String() string {
	return g.Name
}

type dependencies []dependency

type dependency struct {
	Package string `yaml:"package,omitempty"`
	Version string `yaml:"version,omitempty"`
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
	file, err := os.Open(glideFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := fileRead(file)
	if err != nil {
		log.Fatal(err)
	}

	var g glide
	err = yaml.Unmarshal(data, &g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)
}

func fileRead(file *os.File) ([]byte, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, err
}
