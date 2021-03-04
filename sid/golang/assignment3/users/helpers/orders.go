package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetOrders(req *http.Request) (map[int]string, error) {
	var orders map[int]string

	dbReq, err := DuplicateRequestOrders(req)

	if err != nil {
		return orders, err
	}

	client := &http.Client{}
	resp, err := client.Do(dbReq)

	if err != nil {
		return orders, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return orders, err
		}

		return orders, errors.New(string(errMsg))
	}

	ordersJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orders, err
	}

	json.Unmarshal(ordersJSON, &orders)

	return orders, nil
}

func CreateOrder(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequestOrders(req)

	if err != nil {
		return msg, err
	}

	client := &http.Client{}
	resp, err := client.Do(dbReq)

	if err != nil {
		return msg, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return msg, err
		}

		return msg, errors.New(string(errMsg))
	}

	msgBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg, err
	}

	msg = string(msgBytes)

	return msg, err
}

func UpdateOrder(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequestOrders(req)

	if err != nil {
		return msg, err
	}

	client := &http.Client{}
	resp, err := client.Do(dbReq)

	if err != nil {
		return msg, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return msg, err
		}

		return msg, errors.New(string(errMsg))
	}

	msgBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg, err
	}

	msg = string(msgBytes)

	return msg, err
}

func DeleteOrders(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequestOrders(req)

	if err != nil {
		return msg, err
	}

	client := &http.Client{}
	resp, err := client.Do(dbReq)

	if err != nil {
		return msg, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return msg, err
		}

		return msg, errors.New(string(errMsg))
	}

	msgBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg, err
	}

	msg = string(msgBytes)

	return msg, err
}
