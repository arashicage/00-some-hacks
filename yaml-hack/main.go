package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

/*
description: fruits are delicious
fruits:
  apple:
    - red
    - sweet
  lemon:
    - yellow
    - sour
*/

type Fruit struct {
	Name       string
	Properties []string
}

type Config struct {
	Description string
	Fruits      map[string][]string
}

func main() {
	filename, _ := filepath.Abs("./file.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", config.Description)
	fmt.Printf("Value: %#v\n", config.Fruits)
}
