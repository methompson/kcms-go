package usercontroller

// User is a data representation of a user row in the database
type User struct {
	id                 string
	firstName          string
	lastName           string
	username           string
	email              string
	userType           string
	password           string
	passwordResetToken string
	passwordResetDate  int64
	userMeta           string
	dateAdded          int64
	dateUpdated        int64
	enabled            bool
}
