package kcms

import (
	"errors"

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
func MakeKCMS() (*KCMS, error) {
	// fmt.Println("Making KCMS")
	config, err := configuration.ReadConfig()

	if err != nil {
		return nil, err
	}

	var cms *KCMS
	// We will attempt to connect a database and if our result is empty, we will panic

	// Check if the MySQL configuration is empty
	if (config.DB.Mysqldb != configuration.MySQLConfig{}) {
		cms, err = makeMySQLKCMS(config)

		if err != nil {
			return nil, err
		}

		// Check if the MongoDB configuration is empty
	} else if (config.DB.Mongodb != configuration.MongoDBConfig{}) {
		cms, err = makeMongoDBKCMS(config)

		if err != nil {
			return nil, err
		}

		return nil, errors.New("mongoDB no yet implemented")
	}

	if (*cms == KCMS{}) {
		return nil, errors.New("empty KCMS struct created")
	}

	return cms, nil
}

// MakeMySQLKcms will generate a KCMS object with a MySQL database
func makeMySQLKCMS(config configuration.Configuration) (*KCMS, error) {
	// func makeMySQLKcms(dbInstance mysqlcontroller.MySQLCMS) KCMS {
	dbInstance, err := mysqlcontroller.GetMysqlDb(config.DB.Mysqldb)

	if err != nil {
		return nil, err
	}

	cms := KCMS{
		BlogPostController: blogpostcontroller.MySQLBlogPostController{
			Controller: dbInstance,
		},
		UserController: usercontroller.MySQLUserController{
			Controller: dbInstance,
		},
		PageController: pagecontroller.MySQLPageController{
			Controller: dbInstance,
		},
	}

	cms.JWTSecret = config.JWTSecret

	return &cms, nil
}

// MakeMongoDBKcms will generate a KCMS object with a MongoDB database
func makeMongoDBKCMS(config configuration.Configuration) (*KCMS, error) {
	return nil, errors.New("mongoDB no yet implemented")
}
