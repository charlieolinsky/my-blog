package main

/*
	Entry Point for App
	NOTE: Run from Project root using:
	$	air
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

	//Setup Go's ServeMux as Router
	mux := http.NewServeMux()

	//Setup HTTP server routes
	userHandler := handler.NewUserHandler(service.NewUserService(repo.NewUserRepository(db)))
	mux.HandleFunc("/createUser", userHandler.CreateUser)
	mux.HandleFunc("/getUser/{id}", userHandler.GetUserByID)

	//Start HTTP Server
	log.Printf("server running on http://localhost%s", PORT)
	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
