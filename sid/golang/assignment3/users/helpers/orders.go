package helpers

import (
	"net/http"
)

func OrdersHelper(req *http.Request) (*http.Response, error) {
	ordersReq, err := DuplicateRequestOrders(req)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(ordersReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
