package handler

import (
	"net/http"
	"posts/internals/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{services: services}
	return h
}

func (h *Handler) InitRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(h.home))
	r.Handle("/signup", http.HandlerFunc(h.signUp))
	return r
}
