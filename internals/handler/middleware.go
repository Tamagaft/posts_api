package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Token"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(f func(ctx context.Context, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r)
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		c := context.Background()
		c = context.WithValue(c, userCtx, userId)
		f(c, w, r)
	}
}

func getUserId(ctx context.Context) (int, error) {
	id := ctx.Value(userCtx)
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
