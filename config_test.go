package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertConfigVarIsNumber(t *testing.T) {
	filename := "./static/config.json"
	var bytesRead []byte
	bytesRead = Read(filename)
	configMap := ReadJSON(bytesRead)

	memcacheTimeout := configMap["MEMCACHE_TIMEOUT"]
	isNumber, err := IsFloat(memcacheTimeout)

	if isNumber != true {
		t.Error(err)
	}
}

func TestAssertWrongType(t *testing.T) {
	configurationMap := CreateConfigMap([]string{"./static/config.json", "./static/default.json"})
	dbUrl := configurationMap["DATABASE_URL"]
	stringVal, err := IsFloat(dbUrl)
	assert.NotNil(t, err)
	assert.Equal(t, stringVal, false)
}
