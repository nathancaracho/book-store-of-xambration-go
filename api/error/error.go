package error

import "github.com/go-playground/validator/v10"

type ApiFieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ApiError struct {
	Message string          `json:"message"`
	Errors  []ApiFieldError `json:"errors"`
}

func ParseFieldError(err error) ApiError {
	var errors []ApiFieldError
	if validators, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validators {
			errors = append(errors, ApiFieldError{
				fieldError.Field(),
				getErrorMsg(fieldError),
			})
		}
	}
	return ApiError{"Some fields are not correct please fix it.", errors}
}

func getErrorMsg(validator validator.FieldError) string {
	switch validator.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "This field should be lower than" + validator.Param()
	case "gte":
		return "This field should be higher than" + validator.Param()
	}

	return "Some error happen"
}
