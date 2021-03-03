package helpers

import (
	"net/http"
)

const basePath = "http://localhost:2222"

func DuplicateRequest(req *http.Request) (*http.Request, error) {
	queryString := req.URL.Query().Encode()
	dbURL := basePath + req.URL.Path

	if len(queryString) > 0 {
		dbURL += "?" + queryString
	}

	return http.NewRequest(req.Method, dbURL, req.Body)
}
