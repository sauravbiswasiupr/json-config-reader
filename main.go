package main

import (
	"fmt"
	"json-config-reader/config"
)

type message struct {
	Name string
	ID   int64
}

func main() {
	fmt.Println("Readig os...")
	osMap := config.ReadOS()

	for key, value := range osMap {
		fmt.Println("Key: ", key, " Value: ", value)
	}

	// read a file and then parse its contents.
	filename := "./static/config.json"
	var bytesRead []byte
	bytesRead = config.Read(filename)
	fmt.Println("File contents are: ")
	configMap := config.ReadJSON(bytesRead)

	memcacheTimeout := configMap["MEMCACHE_TIMEOUT"]
	fmt.Println("MEMCACHE_TIMEOUT FROM CONFIG: ", memcacheTimeout)

	isInt, err := config.IsFloat(memcacheTimeout)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Number found: ", isInt)
	}

	configurationMap := config.CreateConfigMap([]string{"./static/config.json"})
	fmt.Println("Conf map: ", configurationMap)
}
