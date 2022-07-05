package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func createMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s must be not empty", fe.Field())
	case "email":
		return fmt.Sprintf("%s should be a valid email", fe.Field())
	case "required_without_all":
		return fmt.Sprintf("please fill in one of the %s", fe.Field())
	}
	return fe.Error()
}

func ErrorValidation(err error) []string {
	var errs []string
	ve := err.(validator.ValidationErrors)

	for _, fe := range ve {
		errs = append(errs, createMessage(fe))
	}

	return errs
}
