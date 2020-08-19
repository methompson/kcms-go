package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Starting A Web Server")
	readConfig()
	startServer()
}

func readConfig() {
	dat, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	var config string
	config = string(dat)

	fmt.Println(config)

	var m map[string]interface{}
	err = json.Unmarshal([]byte(config), &m)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range m {
		if key == "mysqldb" {
			getMysqlDb(value)
		}

		fmt.Println(key)
		fmt.Println(value)
	}
}

func getMysqlDb(config interface{}) {
	dbConfig := config.(map[string]interface{})

	fmt.Println(dbConfig["host"])

	for key, value := range dbConfig {
		fmt.Println(key, ": ", value)
	}
}
func getMongoDb() {}

func startServer() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
