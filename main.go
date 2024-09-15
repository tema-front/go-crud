package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tema-front/go-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in environment")
	} else {
		log.Println("PORT has been successfully found")
	}

	// Подключение базы данных
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	} else {
		log.Println("dbURL has been successfully found")
	}

	conn, connErr := sql.Open("postgres", dbURL)
	if connErr != nil {
		log.Fatal("Can't connect to database", connErr)
	} else {
		log.Println("database has been successfully connected")
	}

	if err := conn.Ping(); err != nil {
    log.Fatal("Can't ping the database", err)
	} else {
			log.Println("database has been successfully pinged")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"http://*", "https://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
    AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/create", apiCfg.handlerCreateUser)
	v1Router.Get("/get", apiCfg.handlerGetUser)
	v1Router.Get("/list", apiCfg.handlerGetUsers)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Server starting on port %v", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}