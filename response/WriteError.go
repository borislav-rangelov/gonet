package response

import (
	"net/http"

	"github.com/borislav-rangelov/gonet"
)

// WriteErrors Writes an error response object as json to response
func WriteErrors(code int, errors *map[string][]gonet.Error, w http.ResponseWriter) {
	body := Error{Code: code, Errors: errors}
	WriteJSON(code, body, w)
}
