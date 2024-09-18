package handlers

import (
	"net/http"

	"github.com/tema-front/go-crud/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}