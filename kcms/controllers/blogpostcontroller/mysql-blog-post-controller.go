package blogpostcontroller

import (
	"fmt"

	mysqlcontroller "com.methompson/kcms-go/kcms/db/mysql"
)

// MySQLBlogPostController is an implementation of BlogPostController with MySQL
// implementations for retrieving data
type MySQLBlogPostController struct {
	BaseBlogPostController
	Controller mysqlcontroller.MySQLCMS
}

// GetBlogPostByID gets a BlogPost object from storage using an id
func (inst MySQLBlogPostController) GetBlogPostByID(id string) {
	fmt.Println("Inside MySQL BlogPostController GetBlogPostById")
}

// GetBlogPostBySlug gets a BlogPost object from storage using a slug
func (inst MySQLBlogPostController) GetBlogPostBySlug(slug string) {
	fmt.Println("Inside MySQL BlogPostController GetBlogPostBySlug")
}

// AddBlogPost adds a BlogPost object to storage
func (inst MySQLBlogPostController) AddBlogPost() {
	fmt.Println("Inside MySQL BlogPostController AddBlogPost")
}

// EditBlogPost edits a BlogPost object in storage
func (inst MySQLBlogPostController) EditBlogPost() {
	fmt.Println("Inside MySQL BlogPostController EditBlogPost")
}

// DeleteBlogPost removes a BlogPost object from storage
func (inst MySQLBlogPostController) DeleteBlogPost() {
	fmt.Println("Inside MySQL BlogPostController DeleteBlogPost")
}
