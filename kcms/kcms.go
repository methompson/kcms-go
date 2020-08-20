package kcms

import (
	"com.methompson/go-test/kcms/controllers/blogpostcontroller"
	"com.methompson/go-test/kcms/controllers/pagecontroller"
	"com.methompson/go-test/kcms/controllers/usercontroller"
)

// KCMS represents the cms object for the entire program to use. The blog, page and user controllers
// are created with db-specific logic.
type KCMS struct {
	blogPostController blogpostcontroller.BlogPostController
	pageController     pagecontroller.PageController
	userController     usercontroller.UserController
}

// MakeMySQLKcms will generate a KCMS object with a MySQL database
func MakeMySQLKcms(config interface{}) {
	bpc := blogpostcontroller.MySQLBlogPostController{}

	bpc.AddBlogPost()
	bpc.CheckName()
}

// MakeMongoDBKcms will generate a KCMS object with a MongoDB database
func MakeMongoDBKcms() {}
