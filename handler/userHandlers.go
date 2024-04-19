package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	json.NewEncoder(w).Encode(newUser) //Optionally return the created user object
}
