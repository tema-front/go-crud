package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tema-front/go-crud/config"
	"github.com/tema-front/go-crud/db"
	"github.com/tema-front/go-crud/handlers"
	"github.com/tema-front/go-crud/internal/database"
)


func InitRouter(conn *sql.DB) *chi.Mux {
	apiCfg := database.New(conn)

	router := chi.NewRouter()
	router.Mount("/v1", handlers.NewRouter(apiCfg))

	return router
}

func StartServer(router *chi.Mux, port string) {
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	godotenv.Load(".env")
	config := config.LoadConfig()
	
	conn := db.InitDB(config.DB_URL)

	router := InitRouter(conn)

	StartServer(router, config.PORT)
}
