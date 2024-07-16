package api

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int    `json:"-"`
	ErrorMessage string `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("err[%d]: %s", e.StatusCode, e.ErrorMessage)
}

func NewNotFoundError(apiType, id string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: fmt.Sprintf("%s not found %s", apiType, id),
	}
}
