package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// Read simply loads a json file into memory.
func Read(filename string) []byte {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return raw
}

// ReadJSON reads a json object and spews out the key value pairs.
func ReadJSON(data []byte) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	return m
}

// ReadOS reads out stuff from the operating system.
func ReadOS() map[string]interface{} {
	var osvars []string
	osMap := make(map[string]interface{})

	osvars = os.Environ()

	for _, r := range osvars {
		splits := strings.Split(r, "=")
		key := splits[0]
		value := splits[1]

		osMap[key] = value
	}
	return osMap
}

// CreateConfigMap a map of all variables from config files an OS.env
func CreateConfigMap(filenames []string) map[string]interface{} {
	configMap := make(map[string]interface{})
	for _, filename := range filenames {
		raw := Read(filename)
		m := ReadJSON(raw)
		for k, v := range m {
			configMap[k] = v
		}
	}
	mapOs := ReadOS()
	for k, v := range mapOs {
		configMap[k] = v
	}

	return configMap
}
