package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

// Define all methods related to users
type UserRepository interface {

	//Create User
	CreateUser(ctx context.Context, user model.User) error

	//Read/Get User
	GetUser(ctx context.Context, UserID int) (*model.User, error)

	//Update User
	UpdateUser(ctx context.Context, UserID int, newUser model.User) error

	//Delete User
	DeleteUser(ctx context.Context, UserID int) error
}
