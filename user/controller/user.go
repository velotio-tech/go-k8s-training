package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/velotio-tech/go-k8s-training/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *controller) listUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := c.userService.GetAll(ctx)
	if err != nil {
		log.Println("list users err ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "list users err ", err)
		return
	}
	writeJSON(w, users)
}

func (c *controller) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	user := model.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("inser user err :", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "insert user err :", err)
		return
	}
	user, err = c.userService.Create(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			log.Println("insert user validation failed")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "insert user validation failed")
			return
		}
		log.Println("insert user err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "insert user err :", err)
		return
	}
	writeJSON(w, user)
}

func (c *controller) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	pathParams := mux.Vars(r)
	user := model.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("update user decode err :", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "update user decode err :", err)
		return
	}
	user, err = c.userService.Update(ctx, pathParams["userID"], user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("user not found")
			w.WriteHeader(http.StatusGone)
			fmt.Fprintln(w, "user not found")
			return
		}
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex") {
			log.Println("invalid user id")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid user id")
			return
		}
		log.Println("update user err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "update user err :", err)
		return
	}
	writeJSON(w, user)
}

func (c *controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pathParams := mux.Vars(r)
	err := c.userService.Delete(ctx, pathParams["userID"])
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("user not found")
			w.WriteHeader(http.StatusGone)
			fmt.Fprintln(w, "user not found")
			return
		}
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex") {
			log.Println("invalid user id")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid user id")
			return
		}
		log.Println("insert user err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "insert user err :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
