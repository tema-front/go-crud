package handlers

import (
	"net/http"

	"github.com/tema-front/go-crud/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}