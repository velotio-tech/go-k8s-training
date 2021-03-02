package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
)

func OrderHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	switch req.Method {
	case http.MethodGet:
		getOrderHandler(resp, req, db)
	case http.MethodPost:
		createOrderHandler(resp, req, db)
	case http.MethodPut:
		updateOrderHandler(resp, req, db)
	case http.MethodDelete:
		deleteOrderHandler(resp, req, db)
	default:
		resp.WriteHeader(http.StatusBadRequest)
	}
}

func createOrderHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	type CreateOrderdata struct {
		Details string
		UserID  int
	}

	decoder := json.NewDecoder(req.Body)

	var orderdata CreateOrderdata

	err := decoder.Decode(&orderdata)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	err = helpers.CreateOrder(db, orderdata.UserID, orderdata.Details)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func getOrderHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	var orders map[int]string
	var err error
	var singleRequest bool

	if strings.HasSuffix(req.URL.Path, "/orders") || strings.HasSuffix(req.URL.Path, "/orders/") {
		idInts, err := getIDs(req)

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		if len(idInts) == 0 {
			orders, err = helpers.GetAllOrders(db)
		} else {
			orders, err = helpers.GetOrders(db, idInts)
		}
	} else {
		singleRequest = true
		id, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		orders, err = helpers.GetOrder(db, id)
	}

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 && singleRequest {
		http.NotFound(resp, req)
		return
	}

	ordersJSON, err := json.Marshal(orders)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(ordersJSON))
}

func updateOrderHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	type orderUpdateData struct {
		ID         int
		NewDetails string
	}

	var updatedOrderData orderUpdateData
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&updatedOrderData)
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	err = helpers.UpdateOrder(db, updatedOrderData.ID, updatedOrderData.NewDetails)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func deleteOrderHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	var err error

	if strings.HasSuffix(req.URL.Path, "/orders") || strings.HasSuffix(req.URL.Path, "/orders/") {
		idInts, err := getIDs(req)

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		if len(idInts) == 0 {
			err = helpers.DeleteAllOrders(db)
		} else {
			err = helpers.DeleteOrders(db, idInts)
		}
	} else {
		id, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		err = helpers.DeleteOrder(db, id)
	}

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func getIDs(req *http.Request) ([]int, error) {
	queryParams := req.URL.Query()

	ids, ok := queryParams["ids"]
	var idInts []int

	if !ok {
		return idInts, nil
	} else {
		idInts = make([]int, len(ids))
		for index, id := range ids {
			idInt, err := strconv.Atoi(id)

			if err != nil {
				return idInts, err
			}

			idInts[index] = idInt
		}
		return idInts, nil
	}
}
