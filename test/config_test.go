package config

import (
	"json-config-reader/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertConfigVarIsNumber(t *testing.T) {
	filename := "../static/config.json"
	var bytesRead []byte
	bytesRead = config.Read(filename)
	configMap := config.ReadJSON(bytesRead)

	memcacheTimeout := configMap["MEMCACHE_TIMEOUT"]
	isNumber, err := config.IsFloat(memcacheTimeout)

	if isNumber != true {
		t.Error(err)
	}
}

func TestAssertWrongType(t *testing.T) {
	configurationMap := config.CreateConfigMap([]string{"../static/config.json"})
	dbUrl := configurationMap["DATABASE_URL"]
	stringVal, err := config.IsFloat(dbUrl)
	assert.NotNil(t, err)
	assert.Equal(t, stringVal, false)
}
