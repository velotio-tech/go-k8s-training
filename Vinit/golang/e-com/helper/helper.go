package helper

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func MakeRequest(method string, Url string, body io.ReadCloser) []byte {
	reqUrl, _ := url.Parse(Url)
	req := &http.Request{
		Method: method,
		URL: reqUrl,
		Header: map[string][]string{
			"Content-type": { "application/json; charset=UTF-8" },
		},
		Body: body,
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		//log.Fatal("Error Occurred", err)
		return []byte("{\"StatusCode\" : 404, \"Data\": \"Error Occurred in processing request \"}")
	}
	data, _ := ioutil.ReadAll(res.Body)
	return data
}

func HandleError(err error) bool {
	if err != nil {
		log.Println("Error Occurred")
		return true
	} else {
		return false
	}
}