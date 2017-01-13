package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

func main() {
	yamlData, err := getYaml()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	unformattedJson, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		fmt.Println("Error: Input is not valid YAML")
		os.Exit(1)
	}

	formattedJson, err := prettyPrintJson(unformattedJson)
	if err != nil {
		fmt.Println("Error: Input cannot be converted to JSON")
		os.Exit(1)
	}

	fmt.Printf("%s\n", formattedJson)
}

func getYaml() ([]byte, error) {
	if len(os.Args) > 1 {
		return ioutil.ReadFile(os.Args[1])
	} else {
		return ioutil.ReadAll(os.Stdin)
	}
}

func prettyPrintJson(unformattedJson []byte) ([]byte, error) {
	var jsonData map[string]interface{}

	if err := json.Unmarshal(unformattedJson, &jsonData); err != nil {
		return nil, err
	}

	return json.MarshalIndent(jsonData, "", "  ")
}
