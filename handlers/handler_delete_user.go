package handlers

import (
	"fmt"
	"net/http"

	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)

func (apiCfg ApiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	userID := utils.ParseUserID(w, r)

	err := apiCfg.DB.DeleteUser(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't delete user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}