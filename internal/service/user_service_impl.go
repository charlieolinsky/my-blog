package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/mail"

	"github.com/charlieolinsky/my-blog/internal/model"
	"github.com/charlieolinsky/my-blog/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, newUser model.User) error {

	/* Ensure Valid Email */

	//Email cannot be blank
	if newUser.Email == "" {
		return fmt.Errorf("an email is required")
	}
	//Email must follow standard conventions
	if !isValidEmail(newUser.Email) {
		return fmt.Errorf("invalid email format")
	}
	//Email cannot already be in use
	_, err := s.userRepo.GetUserByEmail(ctx, newUser.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("registration failed, please try again")
	}

	/* Ensure valid Password */

	if len(newUser.Password) < 8 {
		return fmt.Errorf("provided password is too short")
	}

	//Ensure password is hashed
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("password could not be hashed")
	}
	newUser.Password = hashedPassword

	//call the repository
	err = s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
func (s *userService) GetUserByID(ctx context.Context, UserID int) (*model.User, error) {
	//Validate Input
	if UserID <= 0 {
		return nil, fmt.Errorf("invalid user ID: %d", UserID)
	}

	//call the repository
	user, err := s.userRepo.GetUserByID(ctx, UserID)

	//Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	//Email cannot be blank
	if email == "" {
		return nil, fmt.Errorf("an email is required")
	}
	//Email must be of a valid format
	if !isValidEmail(email) {
		return nil, fmt.Errorf("invalid email format")
	}

	//Call repo
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil

}
func (s *userService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	//call the repository
	users, err := s.userRepo.GetAllUsers(ctx)

	//Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}
func (s *userService) UpdateUser(ctx context.Context, UserID int, updatedUser model.User) error {
	//Validate Input
	if UserID <= 0 {
		return fmt.Errorf("invalid user ID: %d", UserID)
	}

	//call the repository
	err := s.userRepo.UpdateUser(ctx, UserID, updatedUser)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil

}
func (s *userService) DeleteUser(ctx context.Context, UserID int) error {
	//Validate Input
	if UserID <= 0 {
		return fmt.Errorf("invalid user ID: %d", UserID)
	}

	//call the repository
	err := s.userRepo.DeleteUser(ctx, UserID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

/* Utility Functions  */

// Check if a given email is valid (using Go's standard library)
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Hash a given password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
