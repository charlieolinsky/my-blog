package repo

import (
	"context"      // Import the context package for managing request-scoped values, cancellation, and deadlines.
	"database/sql" // SQL database interaction functionality.
	"fmt"          // Package for formatted I/O operations.
	"time"         // Package for time-related operations.

	"github.com/charlieolinsky/my-blog/internal/model"
)

// represents a repository for managing user data, encapsulating database interactions.
type userRepository struct {
	db *sql.DB // The database connection pool.
}

// creates and returns an instance of userRepository, which implements the UserRepository interface.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db} // Inject the database connection into the repository.
}

// CreateUser inserts a new user into the database with the provided user information.
func (r *userRepository) CreateUser(ctx context.Context, user model.User) error {
	currentTime := time.Now().UTC() // Capture the current time in UTC for created_at and updated_at fields.

	// SQL query to insert a new user record. Note: Passing nil for deleted_at to indicate the user is active and for updated_at to indicate a user has never been edited.
	query := "INSERT INTO users (role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Role, user.Email, user.Password, user.FirstName, user.LastName, user.ProfilePictureUrl, currentTime, nil, nil)
	return err // Return any error encountered during execution.
}

// GetUser retrieves a user by their ID, including users marked as deleted.
func (r *userRepository) GetUser(ctx context.Context, UserID int) (*model.User, error) {
	var user model.User // Variable to hold the retrieved user data.

	// SQL query to select a user by their ID. Does NOT exclude deleted users as deleted_at is also selected.
	query := "SELECT user_id, role, email, password, first_name, last_name, profilePictureUrl, created_at, updated_at, deleted_at FROM users WHERE user_id=?"
	err := r.db.QueryRowContext(ctx, query, UserID).Scan(&user.UserID, &user.Role, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.ProfilePictureUrl, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetUser Error: No user was found with ID: %d", UserID) // Specific error for not finding a user.
		}
		return nil, err // Return other errors directly.
	}
	return &user, nil // Return the found user.
}

// UpdateUser modifies an existing user's information based on the provided new user data, excluding soft-deleted users.
func (r *userRepository) UpdateUser(ctx context.Context, UserID int, newUser model.User) error {
	currentTime := time.Now().UTC() // Current UTC time for the updated_at field.
	// SQL query to update a user's information, excluding rows where deleted_at is not NULL.
	query := "UPDATE users SET role = ?, email = ?, password = ?, first_name = ?, last_name = ?, profilePictureUrl = ?, updated_at = ? WHERE user_id = ? AND deleted_at IS NULL"
	res, err := r.db.ExecContext(ctx, query, newUser.Role, newUser.Email, newUser.Password, newUser.FirstName, newUser.LastName, newUser.ProfilePictureUrl, currentTime, UserID)

	if err != nil {
		return err // Handle SQL execution errors.
	}

	rowsAffected, err := res.RowsAffected() // Check the number of rows affected by the update.
	if err != nil {
		return err // Handle errors retrieving rows affected.
	}

	if rowsAffected == 0 {
		return fmt.Errorf("UpdateUser Error: No user found with ID: %d", UserID) // Handle case where no rows are updated.
	}

	return nil // Successful update with rows affected.
}

// DeleteUser marks a user as deleted by setting the deleted_at timestamp, instead of removing the user record.
func (r *userRepository) DeleteUser(ctx context.Context, UserID int) error {
	timeNow := time.Now().UTC() // Current UTC time for the deleted_at field.

	// SQL query to soft-delete a user by setting the deleted_at field.
	query := "UPDATE users SET deleted_at=? WHERE user_id=?"
	res, err := r.db.ExecContext(ctx, query, timeNow, UserID)

	if err != nil {
		return err // Handle SQL execution errors.
	}

	rowsAffected, err := res.RowsAffected() // Check the number of rows affected by the update.
	if err != nil {
		return err // Handle errors retrieving rows affected.
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteUser Error: No user found with ID: %d", UserID) // Handle case where no rows are updated.
	}

	return nil // Successful soft-deletion with rows affected.
}
