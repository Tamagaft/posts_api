package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"posts/internals/entity"
)

type userPostPart struct {
	User int `json:"user"`
	Part int `json:"range"`
}

func (h *Handler) CreatePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var p entity.Post
	var userId, err = getUserId(ctx)

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.Text == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.UserPost.CreatePost(userId, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	var postId int
	err := json.NewDecoder(r.Body).Decode(&postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := h.services.GetPostById(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) GetUserPostsRange(w http.ResponseWriter, r *http.Request) {
	var pr userPostPart
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := h.services.GetUserPostsRange(pr.User, pr.Part)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}
