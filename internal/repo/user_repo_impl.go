package repo

import (
	"context"
	"database/sql"
)

/** Represents a pool of database connections. **/
type userRepository struct {
	db *sql.DB 
}

/** Constructor for creating instances of the userRepository **/
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db : db}
}

/** Method on userRepository that creates a user in the DB **/ 
func (r *userRepository) CreateUser(ctx context.Context, user User) error {
	_, err := r.db.ExecContext(ctx, 
		"INSERT INTO users (role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", 
		user.Role, user.Email, user.Password, user.FirstName, user.LastName, user.ProfilePictureUrl, user.CreatedAt, user.UpdatedAt)
	return err
}