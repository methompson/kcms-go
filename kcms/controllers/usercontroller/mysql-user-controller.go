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

// AddUser adds a User object to storage
func (inst MySQLUserController) AddUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	if authToken == nil {
		return "", errors.New("invalid credentials")
	}

	userType := inst.GetUserType(authToken.UserType)

	if !inst.CanAddUser(userType) {
		return "", errors.New("invalid credentials")
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
func (inst MySQLUserController) EditUser(userData InputUserData, authToken *jwtuserdata.JWTUserData) (string, error) {
	// A user MUST be authorized to edit a user
	if authToken == nil {
		return "", errors.New("invalid credentials")
	}

	// We'll get the UserType from the auth token
	editorUserType := inst.GetUserType(authToken.UserType)

	// If the userType doesn't allow editing AND the editor's ID is not the same as
	// the edited user's ID, they can't edit a user
	canEditUser := inst.CanEditUser(editorUserType)
	if !canEditUser && authToken.ID != *userData.ID {
		return "", errors.New("you do not have permission to perform that action")
	}

	// We won't allow a superAdmin to demote themself. We want to make sure at least one
	// superAdmin exists at any given time.
	if authToken.UserType == "superAdmin" &&
		authToken.ID == *userData.ID &&
		userData.UserType != nil &&
		*userData.UserType != "superAdmin" {
		return "", errors.New("cannot demote yourself from superAdmin")
	}

	// We are going to make a slice of queryStrings and a slice of query parameters in
	// empty interfaces.
	queryString := "UPDATE users SET "
	queryParams := make([]interface{}, 0)

	// Changing a password requires that the password be more than 8 characters long.
	// We use bcrypt with a complexity of 12 to encode the password.
	// TODO Allow the configuration options to dictate password length
	if userData.Password != nil {
		if len(*userData.Password) < 8 {
			return "", errors.New("Password must be at least 8 characters long")
		}

		passwordBytes, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), 12)
		if err != nil {
			return "", err
		}
		password := string(passwordBytes)

		queryString += "password = ?, "
		queryParams = append(queryParams, password)
	}

	if userData.FirstName != nil {
		queryString += "firstName = ?, "
		queryParams = append(queryParams, *userData.FirstName)
	}

	if userData.LastName != nil {
		queryString += "lastName = ?, "
		queryParams = append(queryParams, *userData.LastName)
	}

	// Some usernames are invalid, we have to check that
	if userData.Username != nil {
		if len(*userData.Username) < 4 {
			return "", errors.New("username is too short. username must be 4 characters or longer")
		}

		queryString += "username = ?, "
		queryParams = append(queryParams, *userData.Username)
	}

	// We need to make sure that the email is valid
	if userData.Email != nil {
		if !inst.IsEmailValid(*userData.Email) {
			return "", errors.New("email is invalid")
		}

		queryString += "email = ?, "
		queryParams = append(queryParams, *userData.Email)
	}

	// Only an editor is allowed to change a userType
	if userData.UserType != nil {
		if !canEditUser {
			return "", errors.New("not allowed to change userType")
		}

		newUserType := inst.GetUserType(*userData.UserType)

		// SuperAdmins can make new superAdmins, but any rank lower than that cannot make
		// a user the same rank as them. This should enforce a hierarchical ranking structure.
		superRank := ^uint16(0)
		if editorUserType.Rank != superRank && newUserType.Rank >= editorUserType.Rank {
			return "", errors.New("cannot promote user to same or higher userType rank than your own")
		}

		queryString += "userType = ?, "
		queryParams = append(queryParams, *userData.UserType)
	}

	if userData.UserMeta != nil {
		queryString += "userMeta = ?, "
		queryParams = append(queryParams, *userData.UserMeta)
	}

	// Only an editor is allowed to enable or disable a user
	if userData.Enabled != nil {
		if !canEditUser {
			return "", errors.New("not allowed to change enabled")
		}

		queryString += "enabled = ?, "
		queryParams = append(queryParams, *userData.Enabled)
	}

	// dateUpdated will ALWAYS be updated if everything is successful up to now.
	// This will prevent us from having an issue when a user attempts to update
	// their data, but all the data is exactly the same and rowsAffected returns
	// a 0 value. We can now rely on rowsUpdate to return at least 1 for a matched
	// user and 0 when the id doesn't match a user.
	queryString += "dateUpdated = ? "
	queryParams = append(queryParams, time.Now())

	queryString += "WHERE id = ?"
	queryParams = append(queryParams, *userData.ID)

	res, err := inst.Controller.Instance.Exec(queryString, queryParams...)

	// TODO Get potential SQL errors and return custom messages
	// We need to check to determine if there were any SQL errors after executing the query.
	if err != nil {
		returnError := err
		fmt.Println("Error", err)

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

	rows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Rows Affected Error", err)
		return "", err
	}

	// We indicate that no rows were affected.
	if rows < 1 {
		return "", errors.New("user id does not exist")
	}

	return *userData.ID, nil
}

// DeleteUser removes a User object from storage
func (inst MySQLUserController) DeleteUser(id string, authToken *jwtuserdata.JWTUserData) (string, error) {
	// A user MUST be authorized to edit a user
	if authToken == nil {
		return "", errors.New("invalid credentials")
	}

	// We'll get the UserType from the auth token
	editorUserType := inst.GetUserType(authToken.UserType)

	// If the userType doesn't allow editing they can't delete a user.
	canEditUser := inst.CanEditUser(editorUserType)
	if !canEditUser {
		return "", errors.New("you do not have permission to perform that action")
	}

	// superAdmin cannot delete itself
	superRank := ^uint16(0)
	if id == authToken.ID && editorUserType.Rank == superRank {
		return "", errors.New("superAdmin cannot delete itself")
	}

	deletedUser := inst.GetUserByID(id)

	if (deletedUser == User{}) {
		return "", errors.New("user does not exist")
	}

	res, err := inst.Controller.Instance.Exec("DELETE FROM users WHERE id = ? LIMIT 1", id)

	if err != nil {
		returnError := err
		fmt.Println("Error", err)

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

	rows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Rows Affected Error", err)
		return "", err
	}

	// We indicate that no rows were affected.
	if rows < 1 {
		return "", errors.New("user id does not exist")
	}

	return id, nil
}
