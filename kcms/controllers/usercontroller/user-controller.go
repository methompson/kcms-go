package usercontroller

import "com.methompson/kcms-go/kcms/jwtuserdata"

// UserController interface handles all User related tasks
type UserController interface {
	CheckPassword()
	GetUserTypes()
	GetUserRequestToken()
	AuthenticateUserCredentials()

	GetUserByID(id string)
	GetUserByEmail(email string)
	AddUser()
	EditUser()
	DeleteUser()
}

// BaseUserController is a base implementation of the UserController with
// definitions of common functions
type BaseUserController struct{}

// CheckPassword checks a user's password against the hashed version in storage
func (inst BaseUserController) CheckPassword() {}

// GetUserTypes gets all user types
func (inst BaseUserController) GetUserTypes() {}

// GetUserRequestToken extracts the user request token from storage and decodes it
// TODO determine if this is needed anymore
func (inst BaseUserController) GetUserRequestToken() {}

// AuthenticateUserCredentials authenticates the user's request token
func (inst BaseUserController) AuthenticateUserCredentials() {}

// EncodeCredentials will take user credentials and encode them into a JWT
func (inst BaseUserController) EncodeCredentials(jwtuserdata.JWTUserData) string {
	return "123"
}

// GetUserByID gets a User object from storage using an id
func (inst BaseUserController) GetUserByID(id string) {}

// GetUserByEmail gets a User object from storage using an email
func (inst BaseUserController) GetUserByEmail(email string) {}

// AddUser adds a User object to storage
func (inst BaseUserController) AddUser() {}

// EditUser edits a User object in storage
func (inst BaseUserController) EditUser() {}

// DeleteUser removes a User object from storage
func (inst BaseUserController) DeleteUser() {}
