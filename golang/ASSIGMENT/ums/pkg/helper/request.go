package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"ums/pkg/exception"
)

// Author: Arijit Nayak/
// DecodeJSONBody decodes the request body and returns if there in any error of not within
// a exception object
func (m *UserServiceHelper) DecodeJSONBody(r *http.Request, dst interface{}) *exception.Exception {
	err := json.NewDecoder(r.Body).Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			return &exception.Exception{
				Err:        fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset),
				Message:    fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset).Error(),
				StatusText: exception.STATUS_INVALID_DATA,
				StatusCode: http.StatusBadRequest,
			}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return &exception.Exception{
				Err:        fmt.Errorf("request body contains badly-formed JSON"),
				Message:    fmt.Errorf("request body contains badly-formed JSON").Error(),
				StatusText: exception.STATUS_INVALID_DATA,
				StatusCode: http.StatusBadRequest,
			}
		case errors.As(err, &unmarshalTypeError):
			return &exception.Exception{
				Err:        fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset),
				Message:    fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset).Error(),
				StatusText: exception.STATUS_INVALID_DATA,
				StatusCode: http.StatusBadRequest,
			}
		case errors.Is(err, io.EOF):
			return &exception.Exception{
				Err:        fmt.Errorf("request body must not be empty"),
				Message:    fmt.Errorf("request body must not be empty").Error(),
				StatusText: exception.STATUS_INVALID_DATA,
				StatusCode: http.StatusBadRequest,
			}
		default:
			return &exception.Exception{
				Err:        fmt.Errorf("invalid request : unknown error - %v", err),
				Message:    fmt.Errorf("invalid request : unknown error - %v", err).Error(),
				StatusText: exception.STATUS_INVALID_DATA,
				StatusCode: http.StatusBadRequest,
			}
		}
	}
	return nil
}
