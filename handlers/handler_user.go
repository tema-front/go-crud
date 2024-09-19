package handlers

import (
	"github.com/go-chi/chi"
	"github.com/tema-front/go-crud/internal/database"
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