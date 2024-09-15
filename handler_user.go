package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tema-front/go-aggregator/internal/auth"
	"github.com/tema-front/go-aggregator/internal/database"
)

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`;
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, 400, "Name is required")
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
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}	

func (apiCfg apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get users: %v", err))
		return
	}

	parsedUsers := make([]User, len(users))

	for _, v := range users {
		parsedUsers = append(parsedUsers, databaseUserToUser(v))
	}

	respondWithJSON(w, 200, parsedUsers)
}	
