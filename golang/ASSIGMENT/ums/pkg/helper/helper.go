package helper

import (
	"net/http"

	"ums/pkg/exception"
)

// Helper ...
type Helper interface {
	DecodeJSONBody(r *http.Request, dst interface{}) *exception.Exception
	RemoveAllUnNecessarySpaces(str string) string
}

// UserServiceHelper ...
type UserServiceHelper struct{}

// NewUserServiceHelper ...
func NewUserServiceHelper() *UserServiceHelper {
	return &UserServiceHelper{}
}
