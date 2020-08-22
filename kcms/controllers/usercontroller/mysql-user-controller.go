package usercontroller

import (
	"fmt"

	mysqlcontroller "com.methompson/go-test/kcms/db/mysql"
)

// MySQLUserController is an implementation of UserController with MySQL implementations
// for retrieving data
type MySQLUserController struct {
	BaseUserController
	DbInstance mysqlcontroller.MySQLCMS
}

// GetUserByID gets a User object from storage using an id
func (inst MySQLUserController) GetUserByID(id string) {
	fmt.Println("Inside MySQL UserController GetUserByID")
}

// GetUserByEmail gets a User object from storage using an email
func (inst MySQLUserController) GetUserByEmail(email string) {
	fmt.Println("Inside MySQL UserController GetUserByEmail")
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
