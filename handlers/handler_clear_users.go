package handlers

import (
	"fmt"
	"net/http"

	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)

func (apiCfg ApiConfig) handlerClearUsers(w http.ResponseWriter, r *http.Request, _ database.User) {
	err := apiCfg.DB.ClearUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't clear users: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}