package usercontroller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"com.methompson/kcms-go/kcms/jwtuserdata"
	"golang.org/x/crypto/bcrypt"
)

// UserController interface handles all User related tasks
type UserController interface {
	CheckPassword(user User, password string) bool
	GetUserTypes()
	GetUserRequestToken()
	AuthenticateUserCredentials(token string)

	LogUserIn(email *string, username *string, password string) (string, string)
	GetUserByID(id string) User
	GetUserByEmail(email string) User
	GetUserByUsername(username string) User
	AddUser()
	EditUser()
	DeleteUser()
}

// BaseUserController is a base implementation of the UserController with
// definitions of common functions
type BaseUserController struct{}

// CheckPassword checks a user's password against the hashed version in storage
func (inst BaseUserController) CheckPassword(user User, password string) bool {
	fmt.Println("Inside BaseUserController CheckPassword")
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))

	if err == nil {
		return true
	}

	return false
}

// GetUserTypes gets all user types
func (inst BaseUserController) GetUserTypes() {}

// GetUserRequestToken extracts the user request token from storage and decodes it
// TODO determine if this is needed anymore
func (inst BaseUserController) GetUserRequestToken() {}

// AuthenticateUserCredentials authenticates the user's request token
func (inst BaseUserController) AuthenticateUserCredentials(token string) {}

// LogUserIn takes input values from a GraphQL request, gets a user and authenticates the password
// returns a JWT for the user
func (inst BaseUserController) LogUserIn(email *string, username *string, password string) (string, string) {
	return "", ""
}

// EncodeCredentials will take user credentials and encode them into a JWT
func (inst BaseUserController) EncodeCredentials(user User) string {
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

	fmt.Println("jwt stuff")
	fmt.Printf("%+v\n", claims)
	res2B, _ := json.Marshal(claims)
	fmt.Println(string(res2B))
	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println("Error with signing token", err)
		return ""
	}

	fmt.Println(signedToken)
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
func (inst BaseUserController) AddUser() {}

// EditUser edits a User object in storage
func (inst BaseUserController) EditUser() {}

// DeleteUser removes a User object from storage
func (inst BaseUserController) DeleteUser() {}
