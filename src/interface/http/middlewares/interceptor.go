package middlewares

import (
	"encoding/json"
	"net/http"
)

type Interceptor func(http.ResponseWriter, *http.Request) (interface{}, error)

func (fn Interceptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call handler function
	result, err := fn(w, r)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	// Error Handler
	if err != nil {
		handleError(err, w)
		return
	}
	// Response success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return
}

