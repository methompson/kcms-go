package mysqlcontroller

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"com.methompson/go-test/kcms/configuration"
)

// MySQLCMS is a structure that binds all of the different data controllers to a database
type MySQLCMS struct {
	instance *sql.DB
}

func (inst *MySQLCMS) connect(dbInfo map[string]string) {}
func (inst *MySQLCMS) diconnect()                       {}
func (inst *MySQLCMS) checkConnection()                 {}
func (inst *MySQLCMS) validateDbStructure()             {}

// GetMysqlDb takes config file and attempts to connect to a MySQL database
/*
Required variables for a MySQL database include:
- host
- username
- password
- databaseName

The user can specify a port, but if not specified, the default port is 3306
*/
func GetMysqlDb(config configuration.MySQLConfig) MySQLCMS {
	if config.Port == "" {
		fmt.Println("It's empty")
	} else {
		fmt.Println("it's not empty '", config.Host, "'")
	}

	// Checking empty strings
	if config.Host == "" ||
		config.DatabaseName == "" ||
		config.Username == "" ||
		config.Password == "" {
		log.Panic("Invalid DB Parameters")
	}

	port := config.Port
	if port == "" {
		port = "3306"
	}

	mySQLConnectionString := config.Username + ":" + config.Password +
		"@tcp(" + config.Host + ":" + port + ")/" +
		config.DatabaseName

	db, err := sql.Open("mysql", mySQLConnectionString)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	cms := MySQLCMS{
		instance: db,
	}

	results, queryErr := db.Query("SELECT id, name FROM pages")
	if queryErr != nil {
		panic(queryErr)
	}

	fmt.Println("Query Results")
	for results.Next() {
		var name string
		var id int
		err = results.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	return cms
}
