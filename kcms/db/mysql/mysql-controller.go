package mysqlcontroller

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

// MySQLCMS is a structure that binds all of the different data controllers to a database
type MySQLCMS struct {
	instance *sql.DB
}

func (inst *MySQLCMS) connect(dbInfo map[string]string) {}

func (inst *MySQLCMS) diconnect() {}

func (inst *MySQLCMS) checkConnection()     {}
func (inst *MySQLCMS) validateDbStructure() {}

// GetMysqlDb takes config file and attempts to connect to a MySQL database
/*
Required variables for a MySQL database include:
- host
- username
- password
- databaseName

The user can specify a port, but if not specified, the default port is 3306
*/
func GetMysqlDb(config interface{}) MySQLCMS {
	configFile := config.(map[string]interface{})

	dbConfig := make(map[string]string)
	dbConfig["host"] = ""
	dbConfig["username"] = ""
	dbConfig["password"] = ""
	dbConfig["databaseName"] = ""
	dbConfig["port"] = "3306"

	for key, value := range configFile {
		switch vValue := value.(type) {
		case string:
			dbConfig[key] = vValue
		case float64:
			str := strconv.FormatInt(int64(vValue), 10)
			dbConfig[key] = str
		default:
			fmt.Println("Can't identify type")
		}
	}

	if len(dbConfig["host"]) <= 0 ||
		len(dbConfig["username"]) <= 0 ||
		len(dbConfig["password"]) <= 0 ||
		len(dbConfig["databaseName"]) <= 0 ||
		len(dbConfig["port"]) <= 0 {
		log.Panic("Invalid DB Parameters")
	}

	mySQLConnectionString := dbConfig["username"] + ":" + dbConfig["password"] +
		"@tcp(" + dbConfig["host"] + ":" + dbConfig["port"] + ")/" + dbConfig["databaseName"]

	fmt.Println(mySQLConnectionString)

	db, err := sql.Open("mysql", mySQLConnectionString)
	if err != nil {
		panic(err)
	}

	cms := MySQLCMS{
		instance: db,
	}

	fmt.Println("Successful connection")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query, queryErr := db.Query("SELECT id, name FROM pages")
	if queryErr != nil {
		panic(queryErr)
	}

	fmt.Println("Successful Query", query)

	return cms
}
