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
	return &userRepository{db: db}
}

/** Method on userRepository that creates a user in the DB **/
func (r *userRepository) CreateUser(ctx context.Context, user User) error {
	query := "INSERT INTO users (role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Role, user.Email, user.Password, user.FirstName, user.LastName, user.ProfilePictureUrl, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepository) GetUser(ctx context.Context, UserID int) (*User, error) {
	var user User

	query := "SELECT user_id, role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at FROM users WHERE user_id=?"
	err := r.db.QueryRowContext(ctx, query, UserID).Scan(&user.UserID, &user.Role, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.ProfilePictureUrl, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		//Handle the case where the user is not found
		if err == sql.ErrNoRows {
			return nil, nil
		}
		//For other errors, return directly
		return nil, err
	}
	return &user, nil
}
