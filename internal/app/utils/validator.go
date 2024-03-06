package utils

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/go-playground/validator/v10"
)

var lock = &sync.Mutex{}

var validate *validator.Validate

func GetValidator() *validator.Validate {
	if validate == nil {
		lock.Lock()
		defer lock.Unlock()
		if validate == nil {
			validate = initValidator()
		}
	}
	return validate
}

func initValidator() *validator.Validate {
	NewValidate := validator.New(validator.WithRequiredStructEnabled())
	return NewValidate
}

func Validate(r *io.ReadCloser, body interface{}) error {
	reqBody, err := io.ReadAll(*r)
	if err != nil {
		return err
	}

	errUnmarshall := json.Unmarshal(reqBody, &body)
	if errUnmarshall != nil {
		return errUnmarshall
	}

	validator := GetValidator()
	errVal := validator.Struct(body)
	if errVal != nil {
		return errVal
	}

	return nil
}
