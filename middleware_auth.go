package main

import (
	"fmt"
	"net/http"

	"github.com/tema-front/go-aggregator/internal/auth"
	"github.com/tema-front/go-aggregator/internal/database"
)

type authedhandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg apiConfig) middlewareAuth(handler authedhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.AuthByToken(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't auth by apiKey: %v", err))
			return
		}

		handler(w, r, user)
	}
}