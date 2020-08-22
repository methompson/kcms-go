package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"com.methompson/go-test/graph/generated"
	"com.methompson/go-test/graph/model"
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

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
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
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
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

func (r *queryResolver) BlogPosts(ctx context.Context) ([]*model.BlogPost, error) {
	// var emptyStr string = ""

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
