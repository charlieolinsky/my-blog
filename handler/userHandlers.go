package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/charlieolinsky/my-blog/internal/common"
	"github.com/charlieolinsky/my-blog/internal/model"
	"github.com/charlieolinsky/my-blog/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	//Parse the Request Body
	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//Business Logic Execution
	if err := h.userService.CreateUser(r.Context(), newUser); err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	//Response Success
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(newUser) //Optionally return the created user object
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	//Set common headers
	w.Header().Set("Content-Type", "application/json")

	//Parse request URL
	userIDStr := r.PathValue("id")

	//Ensure a userID was provided
	if userIDStr == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}
	//Convert userID string into an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	//Execute Business Logic
	user, err := h.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		}
		return
	}

	//Response Success
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}

}
func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	//Set common headers
	w.Header().Set("Content-Type", "application/json")

	//Parse request URL
	userEmail := r.PathValue("email")

	//Ensure a userID was provided
	if userEmail == "" {
		http.Error(w, "an email is required", http.StatusBadRequest)
		return
	}

	//Ensure email is valid
	if !common.IsValidEmail(userEmail) {
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	}

	//Execute Business Logic
	user, err := h.userService.GetUserByEmail(r.Context(), userEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		}
		return
	}

	//Response Success
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}

}
