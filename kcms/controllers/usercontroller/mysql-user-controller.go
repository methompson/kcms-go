package usercontroller

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	mysqlcontroller "com.methompson/kcms-go/kcms/db/mysql"
	"com.methompson/kcms-go/kcms/jwtuserdata"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// MySQLUserController is an implementation of UserController with MySQL implementations
// for retrieving data
type MySQLUserController struct {
	BaseUserController
	Controller mysqlcontroller.MySQLCMS
}

func (inst MySQLUserController) getUserFromQuery(query string, parameter string) User {
	rows, err := inst.Controller.Instance.Query(query, parameter)

	if err != nil {
		panic(err)
	}

	// inst.ComparePasswordToHash("test", "test")
	var user User

	for rows.Next() {
		err = rows.Scan(&user.id, &user.firstName, &user.lastName, &user.username, &user.email, &user.userType, &user.password, &user.userMeta)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return user
}

// LogUserIn takes input values from a GraphQL request, gets a user and authenticates the password
// returns a JWT for the user
func (inst MySQLUserController) LogUserIn(email *string, username *string, password string, secret string) (string, string) {
	var user User

	// We get a user from
	if email != nil {
		user = inst.GetUserByEmail(*email)
	}

	if username != nil {
		user = inst.GetUserByUsername(*username)
	}

	if (user == User{}) {
		return "", "User Does Not Exist"
	}

	validPass := inst.CheckPassword(user, password)

	if !validPass {
		return "", "Invalid Password"
	}

	token := inst.EncodeCredentials(user, secret)

	return token, ""
}

// GetUserByID gets a User object from storage using an id
func (inst MySQLUserController) GetUserByID(id string) User {
	return inst.getUserFromQuery(`
		SELECT
			id,
			firstName,
			lastName,
			username,
			email,
			userType,
			password,
			userMeta
		FROM users
		WHERE id = ?
		LIMIT 1
	`, id)
}

// GetUserByEmail gets a User object from storage using an email
func (inst MySQLUserController) GetUserByEmail(email string) User {
	return inst.getUserFromQuery(`
    SELECT
      id,
      firstName,
      lastName,
      username,
      email,
      userType,
      password,
      userMeta
    FROM users
    WHERE email = ?
	`, email)
}

// GetUserByUsername gets a User object from storage using an email
func (inst MySQLUserController) GetUserByUsername(username string) User {
	return inst.getUserFromQuery(`
		SELECT
			id,
			firstName,
			lastName,
			username,
			email,
			userType,
			password,
			userMeta
		FROM users
		WHERE username = ?
		LIMIT 1
	`, username)
}

// AddUser adds a User object to storage
func (inst MySQLUserController) AddUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	if authToken == nil {
		return "", errors.New("Invalid Credentials")
	}

	fmt.Println(authToken)
	// fmt.Println(authToken.UserType)

	if !inst.CanAddUser(authToken.UserType) {
		return "", errors.New("Invalid Credentials")
	}

	if len(*userData.Username) < 4 {
		return "", errors.New("username is too short. username must be 4 characters or longer")
	}

	if !inst.IsEmailValid(*userData.Email) {
		return "", errors.New("email is invalid")
	}

	if len(*userData.Password) < 8 {
		return "", errors.New("password is too short. password must be 8 characters or longer")
	}

	query := `
		INSERT INTO users (
			firstName,
			lastName,
			username,
			email,
			userType,
			userMeta,
			password,
			dateAdded,
			dateUpdated,
			enabled
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), 12)

	if err != nil {
		return "", err
	}

	password := string(passwordBytes)
	fmt.Println(password, err)
	now := time.Now()

	res, err := inst.Controller.Instance.Exec(
		query,
		userData.FirstName,
		userData.LastName,
		userData.Username,
		userData.Email,
		userData.UserType,
		userData.UserMeta,
		password,
		now,
		now,
		userData.Enabled,
	)

	// TODO Handle Errors!
	// 1062: Duplicate Entry Error
	if err != nil {
		returnError := err

		driverErr, ok := err.(*mysql.MySQLError)
		if ok {
			code := driverErr.Number
			if code == mysqlerr.ER_DUP_ENTRY {
				returnError = errors.New("duplicate entry. username or email already exists")
			}

			fmt.Println(driverErr.Number, ok)
		}
		return "", returnError
	}

	id, err := res.LastInsertId()

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

// EditUser edits a User object in storage
// TODO Require that a user type in their password to update a user
func (inst MySQLUserController) EditUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	if authToken == nil {
		return "", errors.New("Invalid Credentials")
	}

	if !inst.CanEditUser(authToken.UserType, userData) && authToken.ID != *userData.ID {
		return "", errors.New("Invalid Credentials")
	}

	queryStrings := make([]string, 0)
	queryParams := make([]interface{}, 0)

	if userData.Password != nil {
		if len(*userData.Password) < 8 {
			return "", errors.New("Password must be at least 8 characters long")
		}

		passwordBytes, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), 12)
		if err != nil {
			return "", err
		}
		password := string(passwordBytes)

		queryStrings = append(queryStrings, "password = ?")
		queryParams = append(queryParams, password)
	}

	if userData.FirstName != nil {
		queryStrings = append(queryStrings, "firstName = ?")
		queryParams = append(queryParams, *userData.FirstName)
	}

	if userData.LastName != nil {
		queryStrings = append(queryStrings, "lastName = ?")
		queryParams = append(queryParams, *userData.LastName)
	}

	if userData.Username != nil {
		queryStrings = append(queryStrings, "username = ?")
		queryParams = append(queryParams, *userData.Username)
	}

	if userData.Email != nil {
		queryStrings = append(queryStrings, "email = ?")
		queryParams = append(queryParams, *userData.Email)
	}

	if userData.UserType != nil {
		queryStrings = append(queryStrings, "userType = ?")
		queryParams = append(queryParams, *userData.UserType)
	}

	if userData.UserMeta != nil {
		queryStrings = append(queryStrings, "userMeta = ?")
		queryParams = append(queryParams, *userData.UserMeta)
	}

	if userData.Enabled != nil {
		queryStrings = append(queryStrings, "enabled = ?")
		queryParams = append(queryParams, *userData.Enabled)
	}

	queryString := "UPDATE users SET "

	for i, qs := range queryStrings {
		if i == 0 {
			queryString += qs
		} else {
			queryString += ", " + qs
		}
	}

	queryString += " WHERE id = ?"
	queryParams = append(queryParams, *userData.ID)

	res, err := inst.Controller.Instance.Exec(queryString, queryParams...)

	// TODO Get potential SQL errors and return custom messages
	if err != nil {
		returnError := err
		fmt.Println("Error", err)

		return "", returnError
	}

	rows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Rows Affected Error", err)

		return "", err
	}

	fmt.Println("rows affected", rows)

	// We indicate that no rows were affected.
	if rows < 1 {
		return "", errors.New("no rows were updated")
	}

	return *userData.ID, nil
}

// DeleteUser removes a User object from storage
func (inst MySQLUserController) DeleteUser() {
	fmt.Println("Inside MySQL UserController DeleteUser")
}
