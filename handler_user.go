package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tema-front/go-crud/internal/database"
)

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`;
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if params.Name == "" {
		respondWithError(w, 400, "Name is required")
		return
	}

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse userID: %v", err))
		return
	}

	if userID == (uuid.UUID{}) {
		respondWithError(w, 400, "userID is required")
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}	

func (apiCfg apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	search = strings.TrimSpace(search)

	users, err := apiCfg.DB.GetUsers(r.Context(), search)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get users: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsersToUsers(users))
}	

func (apiCfg apiConfig) handlerClearUsers(w http.ResponseWriter, r *http.Request, _ database.User) {
	err := apiCfg.DB.ClearUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't clear users: %v", err))
		return
	}
	
	respondWithJSON(w, 200, struct{}{})
}

func (apiCfg apiConfig) handlerEditUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse userID: %v", err))
		return
	}

	if userID == (uuid.UUID{}) {
		respondWithError(w, 400, "userID is required")
		return
	}


	updatedUser, err := apiCfg.DB.EditUser(r.Context(), database.EditUserParams{
		ID: userID,
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't edit user: %v", err))
		return
	}

	respondWithJSON(w, 200, updatedUser)
}	

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse userID: %v", err))
		return
	}

	if userID == (uuid.UUID{}) {
		respondWithError(w, 400, "userID is required")
		return
	}

	err = apiCfg.DB.DeleteUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete user: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}	