package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

/** Represents a pool of database connections. **/
type userRepository struct {
	db *sql.DB
}

/** Constructor for creating instances of the userRepository **/
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user User) error {
	currentTime := time.Now().UTC()
	query := "INSERT INTO users (role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Role, user.Email, user.Password, user.FirstName, user.LastName, user.ProfilePictureUrl, currentTime, currentTime)
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

func (r *userRepository) UpdateUser(ctx context.Context, UserID int, newUser User) error {

	currentTime := time.Now().UTC()
	query := "UPDATE users SET role = ?, email = ?, password = ?, first_name = ?, last_name = ?, profilePictureUrl = ?, updated_at = ? WHERE user_id = ?"
	res, err := r.db.ExecContext(ctx, query, newUser.Role, newUser.Email, newUser.Password, newUser.FirstName, newUser.LastName, newUser.ProfilePictureUrl, currentTime, UserID)

	//Handle any SQL execution errors
	if err != nil {
		return err
	}

	//Handle potential errors when checking for rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	//Handle case where no rows are updated
	if rowsAffected == 0 {
		return fmt.Errorf("UpdateUser Error: No user found with ID: %d", UserID)
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, UserID int) error {
	query := "DELETE FROM users WHERE user_id=?"
	res, err := r.db.ExecContext(ctx, query, UserID)

	//Handle SQL Execution Errors
	if err != nil {
		return err
	}
	//Handle potential errors when checking for rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	//Handle case where no rows are updated
	if rowsAffected == 0 {
		return fmt.Errorf("DeleteUser Error: No user found with ID: %d", UserID)
	}

	return nil

}
