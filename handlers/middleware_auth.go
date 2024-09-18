package handlers

import (
	"fmt"
	"net/http"

	"github.com/tema-front/go-crud/internal/auth"
	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)


type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.AuthByToken(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't auth by apiKey: %v", err))
			return
		}

		handler(w, r, user)
	}
}