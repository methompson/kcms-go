package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"com.methompson/kcms-go/graph/generated"
	"com.methompson/kcms-go/graph/model"
	"com.methompson/kcms-go/kcms/headers"
)

func (r *mutationResolver) AddUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	var emptyStr string = ""

	user := &model.User{
		ID:          "",
		FirstName:   &emptyStr,
		LastName:    &emptyStr,
		Username:    "",
		Email:       "",
		UserType:    "",
		UserMeta:    "",
		DateAdded:   1,
		DateUpdaetd: 1,
		Enabled:     true,
	}

	return user, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, id string, input *model.UserInput) (*model.User, error) {
	var emptyStr string = ""

	user := &model.User{
		ID:          "",
		FirstName:   &emptyStr,
		LastName:    &emptyStr,
		Username:    "",
		Email:       "",
		UserType:    "",
		UserMeta:    "",
		DateAdded:   1,
		DateUpdaetd: 1,
		Enabled:     true,
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	return "69", nil
}

func (r *mutationResolver) AddPage(ctx context.Context, input *model.PageInput) (*model.Page, error) {
	page := &model.Page{
		ID:          "1",
		Slug:        "my-slug",
		Enabled:     true,
		Content:     "[]",
		Meta:        "{}",
		DateAdded:   1,
		DateUpdated: 1,
	}

	return page, nil
}

func (r *mutationResolver) EditPage(ctx context.Context, id string, input *model.PageInput) (*model.Page, error) {
	page := &model.Page{
		ID:          "1",
		Slug:        "my-slug",
		Enabled:     true,
		Content:     "[]",
		Meta:        "{}",
		DateAdded:   1,
		DateUpdated: 1,
	}

	return page, nil
}

func (r *mutationResolver) DeletePage(ctx context.Context, id string) (string, error) {
	return "69", nil
}

func (r *mutationResolver) AddBlogPost(ctx context.Context, input *model.BlogPostInput) (*model.BlogPost, error) {
	post := &model.BlogPost{
		ID:          "",
		Name:        "",
		Slug:        "",
		Draft:       true,
		Public:      true,
		Content:     "",
		Meta:        "",
		DateAdded:   1,
		DateUpdated: 1,
	}

	return post, nil
}

func (r *mutationResolver) EditBlogPost(ctx context.Context, id string, input *model.BlogPostInput) (*model.BlogPost, error) {
	post := &model.BlogPost{
		ID:          "",
		Name:        "",
		Slug:        "",
		Draft:       true,
		Public:      true,
		Content:     "",
		Meta:        "",
		DateAdded:   1,
		DateUpdated: 1,
	}

	return post, nil
}

func (r *mutationResolver) DeleteBlogPost(ctx context.Context, id string) (string, error) {
	return "69", nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	return "123", nil
}

func (r *mutationResolver) Signup(ctx context.Context, user model.SignupUser) (string, error) {
	// headers.GetHeaderAuth(ctx)

	// This is a reference
	authToken := headers.GetHeaderAuth(ctx)
	fmt.Println(authToken)
	// fmt.Println(authToken["iat"])
	// claimsType := fmt.Sprintf("%T", authToken)
	// fmt.Println("resolver", claimsType, authToken)
	return "321", nil
}

func (r *queryResolver) Pages(ctx context.Context, pageFilter *model.PageFilter) ([]*model.Page, error) {
	// We have to de-reference the pointer-to-value
	fmt.Println(*pageFilter.ID)
	page := &model.Page{
		ID:          "1",
		Slug:        "my-slug",
		Enabled:     true,
		Content:     "[]",
		Meta:        "{}",
		DateAdded:   1,
		DateUpdated: 1,
	}

	var pages []*model.Page
	pages = append(pages, page)
	return pages, nil
}

func (r *queryResolver) Users(ctx context.Context, userFilter *model.UserFilter) ([]*model.User, error) {
	fmt.Println(userFilter)
	var emptyStr string = ""

	r.KCMS.UserController.GetUserByID("1")

	user := &model.User{
		ID:          "",
		FirstName:   &emptyStr,
		LastName:    &emptyStr,
		Username:    "",
		Email:       "",
		UserType:    "",
		UserMeta:    "",
		DateAdded:   1,
		DateUpdaetd: 1,
		Enabled:     true,
	}

	var users []*model.User
	users = append(users, user)
	return users, nil
}

func (r *queryResolver) BlogPosts(ctx context.Context, blogFilter *model.BlogFilter) ([]*model.BlogPost, error) {
	fmt.Println(blogFilter)

	post := &model.BlogPost{
		ID:          "",
		Name:        "",
		Slug:        "",
		Draft:       true,
		Public:      true,
		Content:     "",
		Meta:        "",
		DateAdded:   1,
		DateUpdated: 1,
	}

	var posts []*model.BlogPost
	posts = append(posts, post)
	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
