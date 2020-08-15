package middlewares

import (
	"encoding/json"
	rest_errors "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/error"
	"net/http"
)

type ErrorHandler func(http.ResponseWriter, *http.Request) error

type Payload struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func (fn ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err != nil {
		switch err.(type) {
		case *rest_errors.RestError:
			err := err.(*rest_errors.RestError)
			w.WriteHeader(err.Status)
			payload := &Payload{
				Type: err.Type,
				Message: err.Message,
			}
			json.NewEncoder(w).Encode(payload)
			return
		default:
			err := rest_errors.NewInternalServerError(`The server failed to handle this request`, err)
			w.WriteHeader(err.Status)
			return
		}
	}
	return
}
