package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	var errorMessage []string
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errorMessage = append(errorMessage, "Invalid email format")
			case "required":
				errorMessage = append(errorMessage, "Field "+err.Field()+" is required")
			case "min":
				if err.Field() == "Password" {
					errorMessage = append(errorMessage, "Field "+err.Field()+" is too short")
				}
			case "eqfield":
				errorMessage = append(errorMessage, "Field "+err.Field()+" is not equal with "+err.Param()+".")
			default:
				errorMessage = append(errorMessage, "Field "+err.Field()+" is invalid")
			}
		}

		return errors.New("Validation Error: " + joinMessage(errorMessage))
	}

	return nil
}

func joinMessage(messages []string) string {
	result := ""
	for i, message := range messages {
		if i > 0 {
			result += ", "
		}
		result += message
	}

	return result
}
