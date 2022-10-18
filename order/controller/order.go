package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/velotio-tech/go-k8s-training/order/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *controller) listUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pathParams := mux.Vars(r)
	orders, err := c.orderService.GetAllByUser(ctx, pathParams["userID"])
	if err != nil {
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex"){
			log.Println("invalid user id :", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid user id :", err)
			return
		}
		log.Println("list orders err ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "list orders err ", err)
		return
	}
	writeJSON(w, orders)
}

func (c *controller) createUserOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	order := model.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		log.Println("inser order err :", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "insert order err :", err)
		return
	}
	order, err = c.orderService.CreateOne(ctx, order)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			log.Println("insert order validation failed")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "insert order validation failed")
			return
		}
		log.Println("insert order err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "insert order err :", err)
		return
	}
	writeJSON(w, order)
}

func (c *controller) updateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	pathParams := mux.Vars(r)
	order := model.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		log.Println("update order decode err :", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "update order decode err :", err)
		return
	}
	order, err = c.orderService.UpdateOne(ctx, pathParams["orderID"], order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("order not found")
			w.WriteHeader(http.StatusGone)
			fmt.Fprintln(w, "order not found")
			return
		}
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex") {
			log.Println("invalid order id")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid order id")
			return
		}
		log.Println("update order err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "update order err :", err)
		return
	}
	writeJSON(w, order)
}

func (c *controller) deleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pathParams := mux.Vars(r)
	err := c.orderService.DeleteOne(ctx, pathParams["orderID"])
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("order not found")
			w.WriteHeader(http.StatusGone)
			fmt.Fprintln(w, "order not found")
			return
		}
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex") {
			log.Println("invalid order id")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid order id")
			return
		}
		log.Println("delete order err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "delete order err :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *controller) deleteUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pathParams := mux.Vars(r)
	err := c.orderService.DeleteAll(ctx, pathParams["userID"])
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("order not found")
			w.WriteHeader(http.StatusGone)
			fmt.Fprintln(w, "order not found")
			return
		}
		if err == primitive.ErrInvalidHex || strings.Contains(err.Error(), "encoding/hex") {
			log.Println("invalid user id")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid user id")
			return
		}
		log.Println("delete orders err :", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "delete orders err :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
