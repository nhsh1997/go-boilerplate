package middlewares

import (
	"encoding/json"
	rest_errors "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/error"
	"net/http"
)

type Payload struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func handleError(err error, w http.ResponseWriter){
	switch err.(type) {
	case *rest_errors.RestError:
		err := err.(*rest_errors.RestError)
		w.WriteHeader(err.Status)
		payload := &Payload{
			Type: err.Type,
			Message: err.Message,
		}
		json.NewEncoder(w).Encode(payload)
	default:
		err := rest_errors.NewInternalServerError(`The server failed to handle this request`, err)
		w.WriteHeader(err.Status)
	}
}
