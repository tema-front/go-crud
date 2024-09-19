package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tema-front/go-crud/db"
	"github.com/tema-front/go-crud/internal/database"
	"github.com/tema-front/go-crud/utils"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewRouter(db *database.Queries) *chi.Mux {
	apiCfg := ApiConfig{DB: db}
	router := chi.NewRouter()
	router.Get("/healthz", HandlerReadiness)
	router.Get("/user/list", apiCfg.handlerGetUsers)
	router.Post("/user/create", apiCfg.handlerCreateUser)
	router.Get("/user/{userID}/get", apiCfg.handlerGetUser)
	router.Put("/user/{userID}/edit", apiCfg.middlewareAuth(apiCfg.handlerEditUser))
	router.Delete("/user/{userID}/delete", apiCfg.middlewareAuth(apiCfg.handlerDeleteUser))
	router.Delete("/user/clear", apiCfg.middlewareAuth(apiCfg.handlerClearUsers))

	return router
}

func (apiCfg ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`;
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if params.Name == "" {
		utils.RespondWithError(w, 400, "Name is required")
		return
	}

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
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
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, db.DatabaseUserToUser(user))
}

func (apiCfg ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userID := utils.ParseUserID(w, r)

	user, err := apiCfg.DB.GetUser(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, db.DatabaseUserToUser(user))
}	

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
		Limit: int32(limit),
		Offset: int32(offset),
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

func (apiCfg ApiConfig) handlerClearUsers(w http.ResponseWriter, r *http.Request, _ database.User) {
	err := apiCfg.DB.ClearUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't clear users: %v", err))
		return
	}
	
	utils.RespondWithJSON(w, 200, struct{}{})
}

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
		ID: userID,
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't edit user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, updatedUser)
}	

func (apiCfg ApiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	userID := utils.ParseUserID(w, r)

	err := apiCfg.DB.DeleteUser(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't delete user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}	