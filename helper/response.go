package helper

import "fmt"

type Response1 struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response2 struct {
	Success bool    `json:"success"`
	Error   *Error2 `json:"error,omitempty"`
}

type Error2 struct {
	Code    int     `json:"code"`
	Message []string `json:"message"`
}

func SuccessfulResponse1(payload interface{}) Response1 {
	return Response1{
		Success: true,
		Payload: payload,
	}
}

func FailedResponse1(code int, message string, payload interface{}) Response1 {
	return Response1{
		Success: false,
		Payload: payload,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}

func FailedResponse2(code int, message []string) Response2 {
	fmt.Println("message :", message)
	return Response2{
		Success: false,
		Error: &Error2{
			Code:    code,
			Message: message,
		},
	}
}
