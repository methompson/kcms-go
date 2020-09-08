package usercontroller

import (
	"errors"
	"testing"

	"com.methompson/kcms-go/graph/model"
)

func TestConvertAddUserInputToInputUserInputsUserCorrectly(t *testing.T) {
	var input model.AddUserInput
	buc := BaseUserController{}
	var output InputUserData

	input = model.AddUserInput{}

	input.Username = "Username"
	input.Email = "Email"
	input.Password = "password"

	output = buc.ConvertAddUserInputToInputUser(input)

	if *output.Username != input.Username ||
		*output.Email != input.Email ||
		*output.Password != input.Password {
		t.Error("Improper conversion to InputUser")
	}

	if *output.FirstName != "" ||
		*output.LastName != "" ||
		*output.UserMeta != "{}" ||
		*output.UserType != "subscriber" ||
		*output.Enabled != true {
		t.Error("Improper conversion for default values")
	}

	input = model.AddUserInput{}

	firstName := "First Name"
	lastName := "Last Name"
	userMeta := "{}"
	userType := "editor"
	enabled := false

	input.FirstName = &firstName
	input.LastName = &lastName
	input.UserMeta = &userMeta
	input.UserType = &userType
	input.Enabled = &enabled

	input.Username = "Username"
	input.Email = "Email"
	input.Password = "password"

	output = buc.ConvertAddUserInputToInputUser(input)

	if *output.Username != input.Username ||
		*output.Email != input.Email ||
		*output.Password != input.Password ||
		*output.FirstName != *input.FirstName ||
		*output.LastName != *input.LastName ||
		*output.UserMeta != *input.UserMeta ||
		*output.UserType != *input.UserType ||
		*output.Enabled != *input.Enabled {
		t.Error("Improper conversion to InputUser")
	}

}

func TestCheckPasswordReturnsTrueWhenBcryptReturnsNil(t *testing.T) {
	old := compareHash
	defer func() { compareHash = old }()

	user := User{}
	buc := BaseUserController{}
	var result bool

	compareHash = func(hashedPassword []byte, unhashedPassword []byte) error {
		return nil
	}

	result = buc.CheckPassword(user, "test")
	if result != true {
		t.Errorf("CheckPassword should have returned true when compareHash returns no error")
	}

	compareHash = func(hashedPassword []byte, unhashedPassword []byte) error {
		return errors.New("Test Error")
	}

	result = buc.CheckPassword(user, "test")
	if result != false {
		t.Errorf("CheckPassword should have returned false when compareHash returns an error")
	}
}
