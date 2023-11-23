package routes

import (
	"net/http"

	"ums/pkg/handlers"

	"github.com/gin-gonic/gin"
)

/*
Route Structure of new routes
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

/*
Routes Array of all available routes
*/
type Routes []Route

// NewRoutes returns all the routes
func NewRoutes(handler handlers.UserHandler) Routes {
	routes := Routes{
		Route{
			Name:        "Health",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: handler.HeathCheck,
		},
		Route{
			Name:        "Create User",
			Method:      http.MethodPost,
			Pattern:     "/users",
			HandlerFunc: handler.CreateUser,
		},
		Route{
			Name:        "Update User",
			Method:      http.MethodPut,
			Pattern:     "/users/:userID",
			HandlerFunc: handler.UpdateUser,
		},
		Route{
			Name:        "Get Users",
			Method:      http.MethodGet,
			Pattern:     "/users",
			HandlerFunc: handler.GetUsers,
		},
		Route{
			Name:        "Get User",
			Method:      http.MethodGet,
			Pattern:     "/users/:userID",
			HandlerFunc: handler.GetUser,
		},
		Route{
			Name:        "Create Order",
			Method:      http.MethodPost,
			Pattern:     "/users/:userID/orders",
			HandlerFunc: handler.CreateOrder,
		},
		Route{
			Name:        "Get Order",
			Method:      http.MethodGet,
			Pattern:     "/users/:userID/orders/:orderID",
			HandlerFunc: handler.GetOrder,
		},
		Route{
			Name:        "Get Orders",
			Method:      http.MethodGet,
			Pattern:     "/users/:userID/orders",
			HandlerFunc: handler.GetOrders,
		},
		Route{
			Name:        "Delete Order",
			Method:      http.MethodDelete,
			Pattern:     "/users/:userID/orders/:orderID",
			HandlerFunc: handler.DeleteOrder,
		},
		Route{
			Name:        "Delete Orders",
			Method:      http.MethodDelete,
			Pattern:     "/users/:userID/orders",
			HandlerFunc: handler.DeleteOrders,
		},
	}
	return routes
}

/*
AttachRoutes Attaches routes to the provided server
*/
func AttachRoutes(server *gin.Engine, routes Routes) {
	for _, route := range routes {
		server.
			Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
