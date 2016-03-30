package configparsetest

import (
	"log"
	"testing"

	"github.com/jarema/configparse"
	"github.com/stretchr/testify/assert"
)

func TestParseConfigFile(t *testing.T) {

	var err error
	expectedResult := map[string]string{
		"testProperty": "testValue",
	}

	resultJSON := make(map[string]string)
	err = configparse.ParseConfigFile(configparse.JSON, "./testcfg.json", &resultJSON)
	if err != nil {
		log.Println("error!", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, resultJSON)

	resultYML := make(map[string]string)
	err = configparse.ParseConfigFile(configparse.YAML, "./testcfg.yml", &resultYML)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, resultYML)
}
