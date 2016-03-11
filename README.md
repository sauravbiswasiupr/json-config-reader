json-config-reader
==================

A package to help load variables from config files and the operating system
environment. In certain cases, one needs to store certain variables in config
files, which are then used in the application dynamically. The package json-config-reader
lets you quickly and easily load all these variables from the config files
and also variables that may have been loaded into the OS environment. It also
allows for type assertions (floats, strings, booleans, ints).

Usage
------

* Install: `go getgithub.com/sauravbiswasiupr/json-config-reader/config`

* import the package config

* create a config map using `map := config.CreateConfigMap(filenames []string)`

* Read a particular value from the map using `value := map[key]`
