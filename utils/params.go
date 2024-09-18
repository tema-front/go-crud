package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func ParseUserID(w http.ResponseWriter, r *http.Request) uuid.UUID {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't parse userID: %v", err))
		return uuid.Nil
	}

	if userID == (uuid.UUID{}) {
		RespondWithError(w, 400, "userID is required")
		return uuid.Nil
	}

	return userID
}