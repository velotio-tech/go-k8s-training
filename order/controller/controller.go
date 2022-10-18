package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/velotio-tech/go-k8s-training/order/service"
	"github.com/velotio-tech/go-k8s-training/order/utils"
)

type controller struct {
	orderService service.Order
}

func NewController() *controller {
	return &controller{
		orderService: service.NewOrderService(),
	}
}

func (c *controller) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{userID}/order", c.listUserOrders).Methods(http.MethodGet)
	r.HandleFunc("/user/{userID}/order", c.deleteUserOrders).Methods(http.MethodDelete)
	r.HandleFunc("/order", c.createUserOrder).Methods(http.MethodPost)
	r.HandleFunc("/order/{orderID}", c.updateOrder).Methods(http.MethodPatch)
	r.HandleFunc("/order/{orderID}", c.deleteOrder).Methods(http.MethodDelete)

	sc := utils.GetServiceConfig()
	if err := http.ListenAndServe(":"+sc.Port, r); err != nil {
		log.Panicln(err)
	}
}
