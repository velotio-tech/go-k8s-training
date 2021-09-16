package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResponseJsonBody(response *http.Response) interface{} {
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var body interface{}
	json.Unmarshal(rawBody, &body)

	return body
}
