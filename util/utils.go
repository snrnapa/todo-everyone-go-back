package util

import "github.com/go-playground/validator/v10"

func ValidationCheck(input interface{}) error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}
