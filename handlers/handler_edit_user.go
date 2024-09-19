package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)

func (apiCfg ApiConfig) handlerEditUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userID := utils.ParseUserID(w, r)

	updatedUser, err := apiCfg.DB.EditUser(r.Context(), database.EditUserParams{
		ID:        userID,
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't edit user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, updatedUser)
}