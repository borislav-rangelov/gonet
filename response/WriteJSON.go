package response

import (
	"encoding/json"
	"net/http"
)

// WriteJSON Sets content type and status code of response and writes the body as json
func WriteJSON(code int, body interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}
