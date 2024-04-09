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
	db, err := sqlite.InitDatabase(dbPath + "my_blog.db")
	if err != nil {
		log.Fatalf("Error Initializing Database -- %v", err)
	}
	defer db.Close() //Close when main() finishes execution

	/* DAL TESTS -- START */
	userRepo := repo.NewUserRepository(db)

	createUser(userRepo, "newemail@new.com")
	getUser(userRepo, 1)
	deleteUser(userRepo, 1)

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
	updateUser(userRepo, updatedUser)

	/* DAL TESTS -- END */
}

// Test Functions for User Repo -- START
func createUser(userRepo repo.UserRepository, email string) {
	newUser := repo.User{
		Role:              "general",
		Email:             email,
		Password:          "password12345",
		FirstName:         "Charles",
		LastName:          "Olinsky",
		ProfilePictureUrl: "",
		CreatedAt:         "",
		UpdatedAt:         "",
	}
	err := userRepo.CreateUser(context.Background(), newUser)
	if err != nil {
		log.Fatalf("DAL Failed to Create User -- %v\n", err)
	}
}
func getUser(userRepo repo.UserRepository, UserID int) {
	user, err := userRepo.GetUser(context.Background(), UserID)
	if err != nil {
		log.Fatalf("DAL Failed to Get User -- %v\n", err)
	} else {
		fmt.Printf("Got User: %s%s%s\n", user.FirstName, " ", user.LastName)
	}
}

func updateUser(userRepo repo.UserRepository, updatedUser repo.User) {
	err := userRepo.UpdateUser(context.Background(), 2, updatedUser)
	if err != nil {
		log.Fatalf("DAL failed to Update User -- %v", err)
	}
}
func deleteUser(userRepo repo.UserRepository, UserID int) {
	err := userRepo.DeleteUser(context.Background(), UserID)
	if err != nil {
		log.Fatalf("DAL Failed to Delete User -- %v\n", err)
	}
}

//Test Functions for User Repo -- END
