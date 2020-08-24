package kcms

import (
	"fmt"
	"log"

	"com.methompson/kcms-go/kcms/configuration"
	"com.methompson/kcms-go/kcms/controllers/blogpostcontroller"
	"com.methompson/kcms-go/kcms/controllers/pagecontroller"
	"com.methompson/kcms-go/kcms/controllers/usercontroller"
	mysqlcontroller "com.methompson/kcms-go/kcms/db/mysql"
)

// KCMS represents the cms object for the entire program to use. The blog, page and user controllers
// are created with db-specific logic.
type KCMS struct {
	BlogPostController blogpostcontroller.BlogPostController
	PageController     pagecontroller.PageController
	UserController     usercontroller.UserController

	JWTSecret string
}

// MakeKCMS will create a KCMS struct and return based upon configuration or panic
func MakeKCMS() KCMS {
	fmt.Println("Making KCMS")
	config := configuration.ReadConfig()

	var cms KCMS
	// We will attempt to connect a database and if our result is empty, we will panic

	// Check if the MySQL configuration is empty
	if (config.DB.Mysqldb != configuration.MySQLConfig{}) {
		dbController := mysqlcontroller.GetMysqlDb(config.DB.Mysqldb)
		cms = makeMySQLKcms(dbController)
		cms.JWTSecret = config.JWTSecret

		// Check if the MongoDB configuration is empty
	} else if (config.DB.Mongodb != configuration.MongoDBConfig{}) {
		fmt.Println("MongoDB Configuration is not empty")
	}

	if (cms == KCMS{}) {
		log.Panic("Empty CMS")
	}

	return cms
}

// MakeMySQLKcms will generate a KCMS object with a MySQL database
func makeMySQLKcms(dbInstance mysqlcontroller.MySQLCMS) KCMS {
	cms := KCMS{
		BlogPostController: blogpostcontroller.MySQLBlogPostController{
			DbInstance: dbInstance,
		},
		UserController: usercontroller.MySQLUserController{
			DbInstance: dbInstance,
		},
		PageController: pagecontroller.MySQLPageController{
			DbInstance: dbInstance,
		},
	}

	return cms
}

// MakeMongoDBKcms will generate a KCMS object with a MongoDB database
func makeMongoDBKcms() {}
