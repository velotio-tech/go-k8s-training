package helpers

import (
	"net/http"
)

const basePathDB = "http://db:2222"
const basePathOrders = "http://orders:2223"

func DuplicateRequest(req *http.Request) (*http.Request, error) {
	queryString := req.URL.Query().Encode()
	dbURL := basePathDB + req.URL.Path

	if len(queryString) > 0 {
		dbURL += "?" + queryString
	}

	return http.NewRequest(req.Method, dbURL, req.Body)
}

func DuplicateRequestOrders(req *http.Request) (*http.Request, error) {
	queryString := req.URL.Query().Encode()
	dbURL := basePathOrders + req.URL.Path

	if len(queryString) > 0 {
		dbURL += "?" + queryString
	}

	return http.NewRequest(req.Method, dbURL, req.Body)
}
