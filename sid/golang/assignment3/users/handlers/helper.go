package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(resp http.ResponseWriter, req *http.Request, err error, code int) {
	fmt.Println("Error handler invoked cuz", err)
	resp.WriteHeader(code)
	resp.Write([]byte(err.Error()))
}
