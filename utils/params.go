package utils

import (
	"fmt"
	"net/http"
	"strconv"

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

func ParsePageAndCount(w http.ResponseWriter, r *http.Request) (limit, offset int32) {
	pageStr := r.URL.Query().Get("page")
	pageInt, pageErr := strconv.Atoi(pageStr)
	if pageErr != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't parse page: %v", pageErr))
		return
	}

	countStr := r.URL.Query().Get("count")
	countInt, countErr := strconv.Atoi(countStr)
	if countErr != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't parse count: %v", countErr))
		return
	}

	limit, offset = int32(countInt), int32((pageInt - 1) * countInt)

	return limit, offset
}
