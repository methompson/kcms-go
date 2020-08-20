package usercontroller

// UserController interface handles all User related tasks
type UserController interface {
	GetUserByID(id string)
	GetUserByEmail(email string)
	AddUser()
	EditUser()
	DeleteUser()
	CheckPassword()
	GetUserTypes()
	GetUserRequestToken()
	AuthenticateUserCredentials()
}

// BaseUserController is a base implementation of the UserController with
// definitions of common functions
type BaseUserController struct{}

func (inst BaseUserController) GetUserByID(id string) {}

func (inst BaseUserController) GetUserByEmail(email string) {}

func (inst BaseUserController) AddUser() {}

func (inst BaseUserController) EditUser() {}

func (inst BaseUserController) DeleteUser() {}

func (inst BaseUserController) CheckPassword() {}

func (inst BaseUserController) GetUserTypes() {}

func (inst BaseUserController) GetUserRequestToken() {}

func (inst BaseUserController) AuthenticateUserCredentials() {}
