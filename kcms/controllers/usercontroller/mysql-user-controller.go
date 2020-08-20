package usercontroller

import "fmt"

// MySQLUserController is an implementation of UserController with MySQL implementations
// for retrieving data
type MySQLUserController struct {
	BaseUserController
}

func (inst MySQLUserController) GetUserByID(id string) {
	fmt.Println("Inside MySQL UserController GetUserByID")
}

func (inst MySQLUserController) GetUserByEmail(email string) {
	fmt.Println("Inside MySQL UserController GetUserByEmail")
}

func (inst MySQLUserController) AddUser() {
	fmt.Println("Inside MySQL UserController AddUser")
}

func (inst MySQLUserController) EditUser() {
	fmt.Println("Inside MySQL UserController EditUser")
}

func (inst MySQLUserController) DeleteUser() {
	fmt.Println("Inside MySQL UserController DeleteUser")
}
