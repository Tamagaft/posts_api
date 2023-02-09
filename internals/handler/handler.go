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
	r.Handle("/signin", http.HandlerFunc(h.signIn))

	r.Handle("/createpost", h.userIdentity(h.CreatePost))
	r.Handle("/getpost", http.HandlerFunc(h.GetPostById))
	r.Handle("/getuserposts", http.HandlerFunc(h.GetUserPostsRange))
	return r
}
