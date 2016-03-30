package configparse

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// constants for file config type
const (
	YAML = "yaml"
	JSON = "json"
)

// ParseConfigFile encodes json or yml config file to given type
func ParseConfigFile(fileType string, filepath string, output interface{}) error {
	var err error
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	switch fileType {
	case YAML:
		err = yaml.Unmarshal(fileContent, output)
		if err != nil {
			return err
		}
	case JSON:
		err = json.Unmarshal(fileContent, output)
	}

	return err
}
