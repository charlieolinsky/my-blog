package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

// Define all methods related to users
type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, UserID int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, UserID int, newUser model.User) error
	DeleteUser(ctx context.Context, UserID int) error
}
