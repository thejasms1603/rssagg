package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/thejasms1603/rssagg/internal/database"

	_ "github.com/lib/pq"
)


type apiConfig struct {
	DB *database.Queries
}


func main() {
	if err:= godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found, relying on system environment variables")
	}

	portString:= os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err:= sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant Connect to the database", err)
	}
	defer conn.Close()
	if err:= conn.Ping(); err != nil {
		log.Fatal("Database not reachable", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))


	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.hanldeFeedFollows))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}	
}