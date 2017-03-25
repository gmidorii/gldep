package main

import (
	"log"
	"os"

	"io/ioutil"

	"encoding/json"

	"bufio"

	"gopkg.in/yaml.v2"
)

type glide struct {
	Name    string       `yaml:"package"`
	Imports dependencies `yaml:"import"`
}

func (g glide) String() string {
	return g.Name
}

type dependencies []*dependency

type dependency struct {
	Package string `yaml:"package,omitempty"`
	Version string `yaml:"version,omitempty"`
}

type manifest struct {
	Dependencies map[string]manifestValue `json:"dependencies,omitempty"`
}

type manifestValue struct {
	Revision string `json:"revision,omitempty"`
}

const glideFile = "glide.yaml"
const manifestFile = "manifest.json"

func main() {
	file, err := os.Open(glideFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := fileRead(file)
	if err != nil {
		log.Fatal(err)
	}

	var g glide
	err = yaml.Unmarshal(data, &g)
	if err != nil {
		log.Fatal(err)
	}

	manifestMap := make(map[string]manifestValue)
	var m manifestValue
	for _, v := range g.Imports {
		m.Revision = v.Version
		manifestMap[v.Package] = m
	}
	mani := manifest{Dependencies: manifestMap}

	json, err := json.Marshal(mani)

	mFile, err := os.Create(manifestFile)
	if err != nil {
		log.Fatal(err)
	}
	defer mFile.Close()

	err = fileWrite(mFile, json)
	if err != nil {
		log.Fatal(err)
	}
}

func fileRead(file *os.File) ([]byte, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, err
}

func fileWrite(file *os.File, json []byte) error {
	writer := bufio.NewWriter(file)
	if _, err := writer.Write(json); err != nil {
		return err
	}
	return writer.Flush()
}
