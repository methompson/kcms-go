package blogpostcontroller

import "fmt"

// BlogPostController Handles all of the tasks associated with dealing with blog data
type BlogPostController interface {
	CanUserModifyBlog()
	CheckSlug()
	CheckName()

	GetBlogPostByID(id string)
	GetBlogPostBySlug(slug string)
	AddBlogPost()
	EditBlogPost()
	DeleteBlogPost()
}

// BaseBlogPostController is a base implementation of the BlogPostController
// with common function definitions
type BaseBlogPostController struct{}

// CanUserModifyBlog checks if a user can modify a blog post
func (inst BaseBlogPostController) CanUserModifyBlog() {}

// CheckSlug checks the validity of a slug, including length and characters
func (inst BaseBlogPostController) CheckSlug() {}

// CheckName checks the validify of a page name, including length and characters
func (inst BaseBlogPostController) CheckName() {
	fmt.Println("Inside BaseBlogPostController CheckName")
}

// GetBlogPostByID gets a BlogPost object from storage using an id
func (inst BaseBlogPostController) GetBlogPostByID(id string) {}

// GetBlogPostBySlug gets a BlogPost object from storage using a slug
func (inst BaseBlogPostController) GetBlogPostBySlug(slug string) {}

// AddBlogPost adds a BlogPost object to storage
func (inst BaseBlogPostController) AddBlogPost() {}

// EditBlogPost edits a BlogPost object in storage
func (inst BaseBlogPostController) EditBlogPost() {}

// DeleteBlogPost removes a BlogPost object from storage
func (inst BaseBlogPostController) DeleteBlogPost() {}
