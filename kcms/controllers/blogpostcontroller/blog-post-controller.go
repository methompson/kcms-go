package blogpostcontroller

import "fmt"

// BlogPostController Handles all of the tasks associated with dealing with blog data
type BlogPostController interface {
	GetBlogPostByID(id string)
	GetBlogPostBySlug(slug string)
	AddBlogPost()
	EditBlogPost()
	DeleteBlogPost()
	CanUserModifyBlog()
	CheckSlug()
	CheckName()
}

// BaseBlogPostController is a base implementation of the BlogPostController
// with common function definitions
type BaseBlogPostController struct{}

func (inst BaseBlogPostController) GetBlogPostByID(id string) {}

func (inst BaseBlogPostController) GetBlogPostBySlug(slug string) {}

func (inst BaseBlogPostController) AddBlogPost() {}

func (inst BaseBlogPostController) EditBlogPost() {}

func (inst BaseBlogPostController) DeleteBlogPost() {}

func (inst BaseBlogPostController) CanUserModifyBlog() {}

func (inst BaseBlogPostController) CheckSlug() {}

func (inst BaseBlogPostController) CheckName() {
	fmt.Println("Inside BaseBlogPostController CheckName")
}
