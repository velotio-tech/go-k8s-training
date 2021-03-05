package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/users/helpers"
)

func OrderHandler(resp http.ResponseWriter, req *http.Request) {
	ordersResponse, err := helpers.OrdersHelper(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	ordersResponseBody, err := ioutil.ReadAll(ordersResponse.Body)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(ordersResponse.StatusCode)
	resp.Write(ordersResponseBody)
}
