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
	inFile, err := os.Open(glideFile)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	glideData, err := readFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	// yaml to struct
	var g glide
	err = yaml.Unmarshal(glideData, &g)
	if err != nil {
		log.Fatal(err)
	}

	// convert
	manifestMap := make(map[string]manifestValue)
	var value manifestValue
	for _, v := range g.Imports {
		value.Revision = v.Version
		manifestMap[v.Package] = value
	}
	manifestData, err := json.MarshalIndent(manifest{Dependencies: manifestMap}, "", "\t")

	outFile, err := os.Create(manifestFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = writeFile(outFile, manifestData)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(file *os.File) ([]byte, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, err
}

func writeFile(file *os.File, json []byte) error {
	writer := bufio.NewWriter(file)
	if _, err := writer.Write(json); err != nil {
		return err
	}
	return writer.Flush()
}
