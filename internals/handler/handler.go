package handler

import (
	"net/http"
	"posts/internals/service"
)

type Handler struct {
	services *service.Service
	Routes   *http.ServeMux
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{services: services, Routes: http.NewServeMux()}
	h.Routes.Handle("/", http.HandlerFunc(home))
	return h
}
