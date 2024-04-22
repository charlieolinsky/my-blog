package service

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type UserService interface {
	CreateUser(ctx context.Context, newUser model.User) error
	GetUserByID(ctx context.Context, UserID int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, UserID int, updatedUser model.User) error
	DeleteUser(ctx context.Context, UserID int) error
}
