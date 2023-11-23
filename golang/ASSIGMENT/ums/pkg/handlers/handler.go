package handlers

import (
	"ums/pkg/domain"
	"ums/pkg/helper"
	"ums/pkg/service"

	"github.com/gin-gonic/gin"
)

// UserHandler ...
type UserHandler interface {
	CreateOrder(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteOrder(c *gin.Context)
	DeleteOrders(c *gin.Context)
	GetOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	HeathCheck(c *gin.Context)
}

// UserHandlerClient handles all user interaction with the outside world
type UserHandlerClient struct {
	domain  domain.User
	hlpr    helper.Helper
	service *service.Service
}

func NewUserHandler(domain domain.User, hlpr helper.Helper,
	service *service.Service,
) *UserHandlerClient {
	return &UserHandlerClient{domain: domain, hlpr: hlpr, service: service}
}
