package validate

import (
	"github.com/docker-golang-mysql/model"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateInputPerson(person *model.Person) error {
	if err := validate.Struct(person); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
