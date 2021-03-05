package handlers

import (
	"net/http"
)

func Hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello"))
	resp.WriteHeader(200)
}
