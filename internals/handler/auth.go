package handler

import (
	"encoding/json"
	"net/http"
	"posts/internals/entity"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Моя домашняя страница!"))
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var u entity.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u.Password == "" || u.Username == "" {
		http.Error(w, "no username or password", http.StatusInternalServerError)
		return
	}

	userId, err := h.services.Authorization.CreateUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userId)
}
