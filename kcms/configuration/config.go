package configuration

import (
	"encoding/json"
	"io/ioutil"
)

// MySQLConfig stores MySQL Configuration information for conencting to the database
type MySQLConfig struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"databaseName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// MongoDBConfig stores MongoDB Configuration information for conencting to the database
type MongoDBConfig struct {
	FullURL  string `json:"fullUrl"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

type dbConfig struct {
	Mysqldb MySQLConfig   `json:"mysqldb"`
	Mongodb MongoDBConfig `json:"mongodb"`
}

// Configuration stores all of the configuration data in an easy to use struct
type Configuration struct {
	DB          dbConfig `json:"db"`
	JWTSecret   string   `json:"jwtSecret"`
	BlogEnabled bool     `json:"blogEnabled"`
}

var ioReadFile = ioutil.ReadFile
var unmarshalJSON = json.Unmarshal

// ReadConfig Reads Configuration Data and returns a Configuration struct
func ReadConfig() (Configuration, error) {
	// This file is read from the perspective of the initially run file,
	// i.e. server.go in the root directory.
	configDat, err := ioReadFile("config.json")
	if err != nil {
		// log.Fatal(err)
		return Configuration{}, err
	}

	// We create a Configuration struct with some default data. Then, we attempt to unmarshal
	// the JSON data into the struct. If the JSON structure is correct, we will have a properly
	// filled struct.
	config := Configuration{
		JWTSecret:   "secret",
		BlogEnabled: true,
	}
	err = unmarshalJSON([]byte(configDat), &config)

	if err != nil {
		// log.Fatal(err)
		return Configuration{}, err
	}

	return config, nil
}
