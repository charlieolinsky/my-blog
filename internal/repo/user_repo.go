package repo

import "context"

type User struct {
	UserID            int
	Role              string
	Email             string
	Password          string
	FirstName         string
	LastName          string
	ProfilePictureUrl string
	CreatedAt         string
	UpdatedAt         string
}

//Define all methods related to users
type UserRepository interface {

	//Create User
	CreateUser(ctx context.Context, user User) error

	//Read/Get User
	GetUser(ctx context.Context, UserID int) (*User, error)

	//Update User
	UpdateUser(ctx context.Context, UserID int, newUser User) error

	//Delete User
	DeleteUser(ctx context.Context, UserID int) error
}
