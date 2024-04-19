package main

/*
	Entry Point for App
	NOTE: Run from Project root using:
		go run ./cmd/main.go
*/

import (
	"log"
	"net/http"
	"os"

	"github.com/charlieolinsky/my-blog/handler"
	"github.com/charlieolinsky/my-blog/internal/repo"
	"github.com/charlieolinsky/my-blog/internal/service"
	"github.com/charlieolinsky/my-blog/pkg/sqlite"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Get Environment Vars
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file -- %v", err)
	}
	DB_PATH := os.Getenv("DB_PATH")
	PORT := os.Getenv("PORT")

	//Initialize Database
	db, err := sqlite.InitDatabase(DB_PATH + "my_blog.db")
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	defer db.Close()

	//Setup HTTP server routes
	userHandler := handler.NewUserHandler(service.NewUserService(repo.NewUserRepository(db)))
	http.HandleFunc("/createUser", userHandler.CreateUser)

	//Start HTTP Server
	log.Printf("server running on http://localhost%s", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
