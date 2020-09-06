package configuration

import (
	"errors"
	"testing"
)

func readFileError(t string) ([]byte, error) {
	return nil, errors.New("Error")
}

func readFileSuccess(t string) ([]byte, error) {
	return []byte("{}"), nil
}

func unmarshalJSONError(input []byte, output interface{}) error {
	return errors.New("Error")
}

func unmarshalJSONSuccess(input []byte, output interface{}) error {
	return nil
}

func TestReadConfigReadfileError(t *testing.T) {
	old := ioReadFile
	defer func() { ioReadFile = old }()

	ioReadFile = readFileError

	_, err := ReadConfig()

	// We expect an error
	if err == nil {
		t.Errorf("When ioutil.ReadFile returns an error, ReadConfig also returns the same error")
	}
}

func TestReadConfigUnmarshalError(t *testing.T) {
	oldRead := ioReadFile
	oldMarsh := unmarshalJSON

	defer func() {
		ioReadFile = oldRead
		unmarshalJSON = oldMarsh
	}()

	ioReadFile = readFileSuccess
	unmarshalJSON = unmarshalJSONError

	_, err := ReadConfig()

	// We expect an error
	if err == nil {
		t.Errorf("ReadConfig should be returning an error")
	}
}

func TestReadConfigReadfileSuccess(t *testing.T) {
	oldRead := ioReadFile
	oldMarsh := unmarshalJSON

	defer func() {
		ioReadFile = oldRead
		unmarshalJSON = oldMarsh
	}()

	ioReadFile = readFileSuccess
	unmarshalJSON = unmarshalJSONSuccess

	config, err := ReadConfig()

	configProto := Configuration{
		JWTSecret:   "secret",
		BlogEnabled: true,
	}

	// We expect an error
	if err != nil {
		t.Errorf("When ioutil.ReadFile doesn't return an error and unsmarshalJSON doesn't return an error, ReadConfig also doesn't return an error")
	}

	if config != configProto {
		t.Errorf("ReadConfig should be returning the exact same type of struct")
	}
}

func TestReadConfigReadfileAddSuccess(t *testing.T) {
	oldRead := ioReadFile
	oldMarsh := unmarshalJSON

	defer func() {
		ioReadFile = oldRead
		unmarshalJSON = oldMarsh
	}()

	ioReadFile = readFileSuccess
	unmarshalJSON = func(input []byte, output interface{}) error {
		outValue := output.(*Configuration)
		outValue.JWTSecret = "ABC"
		outValue.DB = dbConfig{}

		return nil
	}

	config, err := ReadConfig()

	configProto := Configuration{
		JWTSecret:   "ABC",
		BlogEnabled: true,
		DB:          dbConfig{},
	}

	// We expect an error
	if err != nil {
		t.Errorf("When ioutil.ReadFile doesn't return an error and unsmarshalJSON doesn't return an error, ReadConfig also doesn't return an error")
	}

	if config != configProto {
		t.Errorf("ReadConfig should be returning the exact same type of struct")
	}
}
