package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/velotio-tech/go-k8s-training/user/service"
	"github.com/velotio-tech/go-k8s-training/user/utils"
)

type controller struct {
	userService service.User
}

func NewController() *controller {
	return &controller{
		userService: service.NewUserService(),
	}
}

func (c *controller) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/user", c.listUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", c.createUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{userID}", c.updateUser).Methods(http.MethodPatch)
	r.HandleFunc("/user/{userID}", c.deleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/user/{userID}/order", c.getUserOrders).Methods(http.MethodGet)
	r.HandleFunc("/user/{userID}/order", c.deleteUserOrders).Methods(http.MethodDelete)
	r.HandleFunc("/order", c.createUserOrder).Methods(http.MethodPost)
	r.HandleFunc("/order/{orderID}", c.updateOrder).Methods(http.MethodPatch)
	r.HandleFunc("/order/{orderID}", c.deleteOrder).Methods(http.MethodDelete)

	sc := utils.GetServiceConfig()
	if err := http.ListenAndServe(":"+sc.Port, r); err != nil {
		log.Panicln(err)
	}
}
