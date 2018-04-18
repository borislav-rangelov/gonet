package handlers

import (
	"net/http"
)

type ChainFunc func(w http.ResponseWriter, req *http.Request, next *ChainIterator)

type ChainHandler struct {
	chain *[]ChainFunc
}

type ChainIterator struct {
	intex int
	chain *[]ChainFunc
}

func (ex *ChainIterator) next(w http.ResponseWriter, req *http.Request) {
	i := (*ex).intex
	i++
	if i >= len(*ex.chain) {
		return
	}
	(*ex).intex = i
	(*ex.chain)[i](w, req, ex)
}

func (h *ChainHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h.chain == nil || len(*h.chain) == 0 {
		return
	}
	ex := ChainIterator{intex: 0, chain: h.chain}
	(*h.chain)[0](w, req, &ex)
}

func Chain(funcs ...ChainFunc) *ChainHandler {
	return &ChainHandler{chain: &funcs}
}

func LastChain(h http.Handler) ChainFunc {
	return LastChainFunc(h.ServeHTTP)
}

func LastChainFunc(h http.HandlerFunc) ChainFunc {
	return func(w http.ResponseWriter, req *http.Request, next *ChainIterator) {
		h(w, req)
	}
}
