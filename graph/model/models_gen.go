// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BlogFilter struct {
	ID   *string `json:"id"`
	Slug *string `json:"slug"`
}

type BlogPost struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Draft       bool   `json:"draft"`
	Public      bool   `json:"public"`
	Content     string `json:"content"`
	Meta        string `json:"meta"`
	DateAdded   int    `json:"dateAdded"`
	DateUpdated int    `json:"dateUpdated"`
}

type BlogPostInput struct {
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Draft   *bool   `json:"draft"`
	Public  *bool   `json:"public"`
	Content string  `json:"content"`
	Meta    *string `json:"meta"`
}

type Page struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Enabled     bool   `json:"enabled"`
	Content     string `json:"content"`
	Meta        string `json:"meta"`
	DateAdded   int    `json:"dateAdded"`
	DateUpdated int    `json:"dateUpdated"`
}

type PageFilter struct {
	ID   *string   `json:"id"`
	Ids  []*string `json:"ids"`
	Slug *string   `json:"slug"`
}

type PageInput struct {
	Slug    string  `json:"slug"`
	Enabled *bool   `json:"enabled"`
	Content string  `json:"content"`
	Meta    *string `json:"meta"`
}

type SignupUser struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	UserType  *string `json:"userType"`
	Enabled   *bool   `json:"enabled"`
	UserMeta  *string `json:"userMeta"`
}

type User struct {
	ID          string  `json:"id"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	UserType    string  `json:"userType"`
	UserMeta    string  `json:"userMeta"`
	DateAdded   int     `json:"dateAdded"`
	DateUpdaetd int     `json:"dateUpdaetd"`
	Enabled     bool    `json:"enabled"`
}

type UserFilter struct {
	ID       *string `json:"id"`
	Email    *string `json:"email"`
	UserType *string `json:"userType"`
}

type UserInput struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	UserType  *string `json:"userType"`
	UserMeta  *string `json:"userMeta"`
	Enabled   *bool   `json:"enabled"`
	Password  string  `json:"password"`
}
