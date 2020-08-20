package blogpostcontroller

import "fmt"

// MySQLBlogPostController is an implementation of BlogPostController with MySQL
// implementations for retrieving data
type MySQLBlogPostController struct {
	BaseBlogPostController
}

func (inst MySQLBlogPostController) GetBlogPostByID(id string) {
	fmt.Println("Inside MySQL BlogPostController GetBlogPostById")
}

func (inst MySQLBlogPostController) GetBlogPostBySlug(slug string) {
	fmt.Println("Inside MySQL BlogPostController GetBlogPostBySlug")
}

func (inst MySQLBlogPostController) AddBlogPost() {
	fmt.Println("Inside MySQL BlogPostController AddBlogPost")
}

func (inst MySQLBlogPostController) EditBlogPost() {
	fmt.Println("Inside MySQL BlogPostController EditBlogPost")
}

func (inst MySQLBlogPostController) DeleteBlogPost() {
	fmt.Println("Inside MySQL BlogPostController DeleteBlogPost")
}
