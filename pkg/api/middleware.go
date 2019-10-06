package api

import (
	"net/http"
)

type httpResponse struct {
	statusCode int
}

func validateRequest(r *http.Request) (bool, int) {
	// Ensure request is a POST request
	if r.Method != http.MethodPost {
		return true, 415
	}

	// Ensure there is a request body
	if r.ContentLength == 0 {
		return true, 400
	}

	// The mime type of this request has to be application/json
	if r.Header.Get("Content-Type") != "application/json" {
		return true, 415
	}

	return false, 200
}

func postValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hasError, statusCode := validateRequest(r)
		if hasError {
			http.Error(w, http.StatusText(statusCode), statusCode)
			return
		}

		next.ServeHTTP(w, r)
	})
}
