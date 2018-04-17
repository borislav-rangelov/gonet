package handlers

import (
	"net/http"

	"github.com/borislav-rangelov/gonet"
	"github.com/borislav-rangelov/gonet/response"
)

// NotFoundHandler Not found handler
type NotFoundHandler struct {
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := gonet.Error{
		Code:    "not-found",
		Message: "Resource not found.",
		Args:    []interface{}{req.URL.String()}}

	errors := map[string][]gonet.Error{"page": []gonet.Error{err}}
	response.WriteErrors(http.StatusNotFound, &errors, w)
}
