package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/tema-front/go-crud/config"
	"github.com/tema-front/go-crud/db"
	"github.com/tema-front/go-crud/handlers"
	"github.com/tema-front/go-crud/internal/database"
)


func initRouter(conn *sql.DB) *chi.Mux {
	apiCfg := database.New(conn)

	router := chi.NewRouter()
	router.Mount("/v1", handlers.NewRouter(apiCfg))

	return router
}

func startServer(router *chi.Mux, port string) {
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Couldn't start server %v", err)
	}

	log.Printf("Server starting on port %v", port)
}

func main() {
	config.LoadEnv()
	config := config.LoadConfig()
	
	conn := db.InitDB(config.DB_URL)

	router := initRouter(conn)

	startServer(router, config.PORT)
}
