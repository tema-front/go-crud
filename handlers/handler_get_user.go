package handlers

import (
	"fmt"
	"net/http"

	"github.com/tema-front/go-crud/db"
	"github.com/tema-front/go-crud/utils"
)

func (apiCfg ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userID := utils.ParseUserID(w, r)

	user, err := apiCfg.DB.GetUser(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, db.DatabaseUserToUser(user))
}