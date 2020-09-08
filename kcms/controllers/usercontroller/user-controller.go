package usercontroller

import (
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"

	"com.methompson/kcms-go/graph/model"
	"com.methompson/kcms-go/kcms/jwtuserdata"
	"golang.org/x/crypto/bcrypt"
)

// UserController interface handles all User related tasks
type UserController interface {
	CheckPassword(user User, password string) bool
	GetUserTypes() map[string]UserType
	GetUserType(userType string) UserType
	CanAddUser(userType UserType) bool
	CanEditUser(userType UserType) bool
	GetUserRequestToken()
	IsEmailValid(email string) bool
	ConvertAddUserInputToInputUser(input model.AddUserInput) InputUserData
	ConvertEditUserInputToInputUser(input model.EditUserInput) InputUserData

	LogUserIn(email *string, username *string, password string, secret string) (string, string)
	GetUserByID(id string) User
	GetUserByEmail(email string) User
	GetUserByUsername(username string) User
	AddUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error)
	EditUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error)
	DeleteUser(id string, authToken *jwtuserdata.JWTUserData) (string, error)
}

// UserType is a convenience structure to easily find and determine the permissions of a user type
type UserType struct {
	UserType    string
	Permissions []string
	Rank        uint16
}

// InputUserData struct provides the ability to add all necessary new user data into a single structure
type InputUserData struct {
	ID        *string
	FirstName *string
	LastName  *string
	Username  *string
	Email     *string
	UserType  *string
	UserMeta  *string
	Enabled   *bool
	Password  *string
}

// BaseUserController is a base implementation of the UserController with
// definitions of common functions
type BaseUserController struct{}

// ConvertAddUserInputToInputUser will convert the GraphQL AddUserInput model into an
// InputUserData struct for easy insertion into the database. This function
// checks for nil input and pre-fills with defaults.
func (inst BaseUserController) ConvertAddUserInputToInputUser(input model.AddUserInput) InputUserData {
	userData := InputUserData{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  &input.Username,
		Email:     &input.Email,
		UserType:  input.UserType,
		UserMeta:  input.UserMeta,
		Enabled:   input.Enabled,
		Password:  &input.Password,
	}

	emptyStr := ""

	if input.FirstName == nil {
		userData.FirstName = &emptyStr
	}
	if input.LastName == nil {
		userData.LastName = &emptyStr
	}

	subscriber := "subscriber"
	if input.UserType == nil {
		userData.UserType = &subscriber
	}

	meta := "{}"
	if input.UserMeta == nil {
		userData.UserMeta = &meta
	}

	enabled := true
	if input.Enabled == nil {
		userData.Enabled = &enabled
	}

	return userData
}

// ConvertEditUserInputToInputUser will convert the GraphQL EditUserInput model into
// an InputUserData structu for easy insertion into the database. This function does not
// check for nil input, because the user is allowed to not send all data.
func (inst BaseUserController) ConvertEditUserInputToInputUser(input model.EditUserInput) InputUserData {
	return InputUserData{
		ID:        &input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Email:     input.Email,
		UserType:  input.UserType,
		UserMeta:  input.UserMeta,
		Enabled:   input.Enabled,
		Password:  input.Password,
	}
}

var compareHash = bcrypt.CompareHashAndPassword

// CheckPassword checks a user's password against the hashed version in storage
func (inst BaseUserController) CheckPassword(user User, password string) bool {
	err := compareHash([]byte(user.password), []byte(password))

	if err == nil {
		return true
	}

	return false
}

// GetUserTypes gets all user types
func (inst BaseUserController) GetUserTypes() map[string]UserType {
	userTypes := make(map[string]UserType)

	userTypes["superAdmin"] = UserType{
		UserType:    "superAdmin",
		Permissions: []string{"view", "edit"},
		Rank:        ^uint16(0),
	}
	userTypes["admin"] = UserType{
		UserType:    "admin",
		Permissions: []string{"view", "edit"},
		Rank:        ^uint16(0) - 1,
	}
	userTypes["editor"] = UserType{
		UserType:    "editor",
		Permissions: []string{"view"},
		Rank:        32767, // All 1s except the furthest left binary digit
	}
	userTypes["subscriber"] = UserType{
		UserType:    "subscriber",
		Permissions: []string{},
		Rank:        0, // Cannot do ANYTHING
	}

	return userTypes
}

// GetUserType gets the usertype struct from a string name of the User Type
func (inst BaseUserController) GetUserType(userType string) UserType {
	// If the userType doesn't actually exist, we get a zeroed version of a user type that has no permissions and a 0 rank
	return inst.GetUserTypes()[userType]
}

// CanAddUser provides a boolean indicating whether the current user is allowed
// to add or edit other users.
func (inst BaseUserController) CanAddUser(userType UserType) bool {
	// t := inst.GetUserTypes()[userType]
	result := false

	for _, s := range userType.Permissions {
		if s == "edit" {
			result = true
		}
	}

	return result
}

// CanEditUser provides a boolean indicating whether the current user is allowed
// to add or edit other users.
func (inst BaseUserController) CanEditUser(userType UserType) bool {
	if !inst.CanAddUser(userType) {
		return false
	}
	result := true

	// Can't move from superAdmin to anything else.

	return result
}

// GetUserRequestToken extracts the user request token from storage and decodes it
// TODO determine if this is needed anymore
func (inst BaseUserController) GetUserRequestToken() {}

// IsEmailValid is a regex function that checks the validity of an email address.
// taken from https://golangcode.com/validate-an-email-address/
func (inst BaseUserController) IsEmailValid(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 || len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}

// LogUserIn takes input values from a GraphQL request, gets a user and authenticates the password
// returns a JWT for the user
func (inst BaseUserController) LogUserIn(email *string, username *string, password string, secret string) (string, string) {
	return "", ""
}

// EncodeCredentials will take user credentials and encode them into a JWT
func (inst BaseUserController) EncodeCredentials(user User, secret string) string {
	// Expiration time will be 4 hours from now.
	exp := time.Now().Unix() + 4*60*60
	claims := jwtuserdata.JWTUserData{
		ID:        user.id,
		FirstName: user.firstName,
		LastName:  user.lastName,
		Username:  user.username,
		Email:     user.email,
		UserType:  user.userType,
		UserMeta:  user.userMeta,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return signedToken
}

// GetUserByID gets a User object from storage using an id
func (inst BaseUserController) GetUserByID(id string) User {
	return User{}
}

// GetUserByEmail gets a User object from storage using an email
func (inst BaseUserController) GetUserByEmail(email string) User {
	return User{}
}

// GetUserByUsername gets a User object from storage using a username
func (inst BaseUserController) GetUserByUsername(username string) User {
	return User{}
}

// AddUser adds a User object to storage
func (inst BaseUserController) AddUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	return "", nil
}

// EditUser edits a User object in storage
func (inst BaseUserController) EditUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	return "", nil
}

// DeleteUser removes a User object from storage
func (inst BaseUserController) DeleteUser(id string, authToken *jwtuserdata.JWTUserData) (string, error) {
	return "", nil
}
