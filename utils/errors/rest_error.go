package errors

import "net/http"

//RestErr a..
type RestErr struct{
	Message string	`json:"message"`
	Status	int		`json:"status"`
	Error 	string	`json:"error"`
}

// NewBadRequestError a...
func NewBadRequestError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_fount",
	}
}

func NewInternalServerError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "Internal_server_error",
	}
}