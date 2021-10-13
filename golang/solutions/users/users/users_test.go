package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	URL = "http://localhost:8080/"
)

func TestCreateUser(t *testing.T) {
	u := user{
		Name:     "user1",
		Location: "Pune",
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(u)
	res, err := http.Post(URL+"users", "Content-Type: application/json", payload)
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("Cannot create user.")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Body : %s", body)
}

func TestGetUser(t *testing.T) {
	expectedUser := user{
		Name:     "user1",
		Location: "Pune",
	}

	res, err := http.Get(URL + "users/" + expectedUser.Name)
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get user.")
	}
	defer res.Body.Close()
	var u user
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	if expectedUser.Name != u.Name || expectedUser.Location != u.Location {
		t.Fatal("Invalid user returned.")
	}
}

func TestGetAllUsers(t *testing.T) {
	expectedUser := user{
		Name:     "user1",
		Location: "Pune",
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(expectedUser)
	res, err := http.Get(URL + "users")
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get user.")
	}
	defer res.Body.Close()
	var u usersMap
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	if len(u) != 1 ||
		expectedUser.Name != u[expectedUser.Name].Name ||
		expectedUser.Location != u[expectedUser.Name].Location {
		t.Fatal("Invalid user returned.")
	}
}

func TestDeleteUser(t *testing.T) {
	expectedUser := user{
		Name:     "user1",
		Location: "Pune",
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", URL+"users/"+expectedUser.Name, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot delete user.")
	}
}

func TestGetAllUsersEmpty(t *testing.T) {
	expectedUser := user{
		Name:     "user1",
		Location: "Pune",
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(expectedUser)
	res, err := http.Get(URL + "users")
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get user.")
	}
	defer res.Body.Close()
	var u usersMap
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	if len(u) != 0 {
		t.Fatal("No users expected but found some.")
	}
}
