package main

/*
	Entry Point for App
	NOTE: Run from Project root using:
		go run ./cmd/main.go
*/

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charlieolinsky/my-blog/internal/repo"
	"github.com/charlieolinsky/my-blog/pkg/sqlite"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Load vars from .env into the application environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file -- %v", err)
	}

	//Get db file path from env
	dbPath := os.Getenv("DB_PATH")

	//Initialize Database
	db, err := sqlite.InitDatabase(dbPath + "test.db")
	if err != nil {
		log.Fatalf("Error Initializing Database -- %v", err)
	}
	defer db.Close() //Close when main() finishes execution

	/* DAL TESTS */
	userRepo := repo.NewUserRepository(db)

	//Create User

	// newUser := repo.User{
	// 	Role:              "admin",
	// 	Email:             "newuser4@blog.com",
	// 	Password:          "hey123456",
	// 	FirstName:         "Charles",
	// 	LastName:          "Olinsky",
	// 	ProfilePictureUrl: "",
	// 	CreatedAt:         "now",
	// 	UpdatedAt:         "now",
	// }

	// err = userRepo.CreateUser(context.Background(), newUser)
	// if err != nil {
	// 	log.Fatalf("DAL Failed to Create User -- %v\n", err)
	// }

	//Get User
	user, err := userRepo.GetUser(context.Background(), 1)
	if err != nil {
		log.Fatalf("DAL Failed to Get User -- %v\n", err)
	} else {
		fmt.Printf("Got User: %s%s%s\n", user.FirstName, " ", user.LastName)
	}

	//Update User
	updatedUser := repo.User{
		Role:              "general",
		Email:             "newuser26@blog.com",
		Password:          "hey1233",
		FirstName:         "Charles",
		LastName:          "Olinsky",
		ProfilePictureUrl: "",
		CreatedAt:         "",
		UpdatedAt:         "",
	}
	err = userRepo.UpdateUser(context.Background(), 2, updatedUser)
	if err != nil {
		log.Fatalf("DAL failed to Update User -- %v", err)
	}

	//Delete User
	err = userRepo.DeleteUser(context.Background(), 3)
	if err != nil {
		log.Fatalf("DAL Failed to Delete User -- %v\n", err)
	}

}
