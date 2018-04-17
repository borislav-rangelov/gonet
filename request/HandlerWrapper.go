package request

import (
	"net/http"
)

// HandlerWrapper Not found handler
type HandlerWrapper struct {
	Handler func(w http.ResponseWriter, req *http.Request)
}

func (h *HandlerWrapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Handler(w, req)
}
