package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetUsers(req *http.Request) (map[int]string, error) {
	var users map[int]string

	dbReq, err := DuplicateRequest(req)

	if err != nil {
		return users, err
	}

	client := &http.Client{}
	resp, err := client.Do(dbReq)

	if err != nil {
		return users, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return users, err
		}

		return users, errors.New(string(errMsg))
	}

	usersJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return users, err
	}

	json.Unmarshal(usersJSON, &users)

	return users, nil
}

func CreateUser(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequest(req)

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

func UpdateUser(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequest(req)

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

func DeleteUsers(req *http.Request) (string, error) {
	var msg string
	var err error

	dbReq, err := DuplicateRequest(req)

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
