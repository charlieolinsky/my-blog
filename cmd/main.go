package main

/*
	Entry Point for App
	NOTE: Run from Project root using:
		go run ./cmd/main.go
*/

import (
	"log"
	"os"

	"github.com/charlieolinsky/my-blog/pkg/sqlite"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	//Load vars from .env into the application environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file -- %v", err)
	}

	//Get db file path from env
	dbPath := os.Getenv("DB_PATH")

	//Initialize Database
	db, err := sqlite.InitDataBase(dbPath)
	if err != nil {
		log.Fatalf("Error Initializing Database -- %v", err)
	}
	defer db.Close() //Close when main() finishes execution

}
