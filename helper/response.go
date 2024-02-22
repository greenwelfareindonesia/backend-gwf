package helper

type Response1 struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
