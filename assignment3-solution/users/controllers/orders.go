package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	orderServiceHost     = "localhost"
	orderServicePort     = "9091"
	orderServiceEndpoint = "http://" + orderServiceHost + ":" + orderServicePort
)

type Order struct {
	OrderId   int64  `json:"order_id"`
	OrderName string `json:"order_name"`
	Price     int64  `json:"price"`
}

var newOrder Order

func sendRequestToOrderService(context *gin.Context, httpMethod string, url string) {
	fmt.Println("In the proxy, url = ", url)

	client := http.Client{}

	req, err := http.NewRequestWithContext(context, httpMethod, url, context.Request.Body)
	if err != nil {
		fmt.Println("Error occurred while new Request!", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during bindjson", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during reading body", err)
	}

	if httpMethod == "PUT" {
		context.IndentedJSON(http.StatusOK, "Row updated!")
		return
	}
	err = json.Unmarshal(bodyBytes, &newOrder)

	if err != nil {
		fmt.Println("Error during json unmarshalling", err)
	}

	fmt.Println("resp code:", resp.StatusCode, ", neworder: ", newOrder)
	context.IndentedJSON(http.StatusOK, string(bodyBytes))
}

func GetAllOrders(context *gin.Context) {
	uname := context.Param("username")
	fmt.Println("uname: ", uname)

	// Generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders"
		sendRequestToOrderService(context, "GET", url)

	}

}

func GetOrderByOrderId(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")
	fmt.Println("uname: ", uname, " oid:", oid)

	// generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders/" + oid
		sendRequestToOrderService(context, "GET", url)
	}
}

func DeleteAllOrders(context *gin.Context) {
	uname := context.Param("username")
	fmt.Println("uname: ", uname)

	// generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders"
		sendRequestToOrderService(context, "DELETE", url)

	}
}

func DeleteOrder(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")
	fmt.Println("uname: ", uname)

	// generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders/" + oid
		sendRequestToOrderService(context, "DELETE", url)

	}
}

func UpdateOrder(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")
	fmt.Println("uname: ", uname, " order_id:", oid)

	// generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders/" + oid
		sendRequestToOrderService(context, "PUT", url)

	}
}

func CreateOrders(context *gin.Context) {
	uname := context.Param("username")
	fmt.Println("uname: ", uname)

	// generate a request for order service
	// before this check whether the user exists
	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("User not found!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		url := orderServiceEndpoint + "/user/" + uname + "/orders"
		sendRequestToOrderService(context, "POST", url)

	}
}
