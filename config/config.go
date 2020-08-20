package config

import (
	"encoding/json"
	"fmt"
	"log"

	"com.methompson/go-test/kcms"
)

// ReadConfig Reads Configuration Data
func ReadConfig(dat []byte) {
	var config string
	config = string(dat)

	fmt.Println(config)

	var m map[string]interface{}
	err := json.Unmarshal([]byte(config), &m)

	if err != nil {
		log.Fatal(err)
	}

	var cms kcms.KCMS

	for key, value := range m {
		if key == "mysqldb" {
			// cms := kcms.MakeMySQLKcms(value)
			kcms.MakeMySQLKcms(value)
			fmt.Println(cms)
			fmt.Println(value)
		}
	}

	// if cms == nil {
	// 	log.Fatal("Invalid database configuration")
	// }
}

func getMongoDb() {}
