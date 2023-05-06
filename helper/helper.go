package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func APIresponse(status int, data interface{}) Response {
	response := Response{
		Status: status,
		Data:   data,
	}
	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
