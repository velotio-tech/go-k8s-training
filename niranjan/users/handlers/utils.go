package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ToJSON returns converts data to json and writes to ResponseWriter stream
func ToJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

// CheckError return fatal if present
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// BodyParser converts the json data in slice and returns it
func BodyParser(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}
