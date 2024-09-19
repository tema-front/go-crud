package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tema-front/go-crud/db"
	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)

func (apiCfg ApiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	search = strings.TrimSpace(search)

	limit, offset := utils.ParsePageAndCount(w, r)

	order := r.URL.Query().Get("order")
	if order != "ASC" && order != "DESC" && order != "" {
		utils.RespondWithError(w, 400, "Error in order param")
		return
	}

	users, err := apiCfg.DB.GetUsers(r.Context(), database.GetUsersParams{
		Column1: search,
		Column2: order,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get users: %v", err))
		return
	}

	if order == "DESC" {
		users = utils.GetReversedSlice(users)
	}

	utils.RespondWithJSON(w, 200, db.DatabaseUsersToUsers(users))
}