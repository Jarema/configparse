// Package configparse is an utility to parse json or yaml config files straight
// into golang type of choice while saving few lines of codes.
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
// you need to add json/yaml tags if you're passing struct
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
