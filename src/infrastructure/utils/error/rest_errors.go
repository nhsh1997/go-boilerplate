package rest_errors

import (
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Type string `json:"type"`
	Causes []interface{} `json:"causes"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestError{
	return &RestError{
		Message: message,
		Status:    http.StatusBadRequest,
		Type:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestError{
	return &RestError{
		Message: message,
		Status:    http.StatusNotFound,
		Type:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestError{
	result := &RestError{
		Message: message,
		Status:    http.StatusInternalServerError,
		Type:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}