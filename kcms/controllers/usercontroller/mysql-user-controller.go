package usercontroller

import (
	"fmt"
	"log"

	mysqlcontroller "com.methompson/kcms-go/kcms/db/mysql"
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

	fmt.Println("Query Results")
	// inst.ComparePasswordToHash("test", "test")
	var user User

	for rows.Next() {
		err = rows.Scan(&user.id, &user.firstName, &user.lastName, &user.username, &user.email, &user.userType, &user.password, &user.userMeta)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return user
}

// LogUserIn takes input values from a GraphQL request, gets a user and authenticates the password
// returns a JWT for the user
func (inst MySQLUserController) LogUserIn(email *string, username *string, password string) (string, string) {
	var user User

	// We get a user from
	if email != nil {
		fmt.Println("Login w/ Email", *email, password)
		user = inst.GetUserByEmail(*email)
	}

	if username != nil {
		fmt.Println("Login w/ Username", *username, password)
		user = inst.GetUserByUsername(*username)
	}

	if (user == User{}) {
		return "", "User Does Not Exist"
	}

	validPass := inst.CheckPassword(user, password)

	fmt.Println("Valid Password?", validPass)

	if !validPass {
		return "", "Invalid Password"
	}

	// inst.EncodeCredentials(user)
	token := inst.EncodeCredentials(user)

	return token, ""
}

// GetUserByID gets a User object from storage using an id
func (inst MySQLUserController) GetUserByID(id string) User {
	fmt.Println("Inside MySQL UserController GetUserByID")

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
	fmt.Println("Inside MySQL UserController GetUserByEmail")

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
	fmt.Println("Inside MySQL UserController GetUserByUsername")

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
func (inst MySQLUserController) AddUser() {
	fmt.Println("Inside MySQL UserController AddUser")
}

// EditUser edits a User object in storage
func (inst MySQLUserController) EditUser() {
	fmt.Println("Inside MySQL UserController EditUser")
}

// DeleteUser removes a User object from storage
func (inst MySQLUserController) DeleteUser() {
	fmt.Println("Inside MySQL UserController DeleteUser")
}
