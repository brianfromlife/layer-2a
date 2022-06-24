package query

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// ReadAllUsersHandler class for all required interfaces to read all users
type ReadAllUsersHandler struct {
	repo user.IRepository
}

// NewReadAllUsersHandler create new instance
func NewReadAllUsersHandler(repository user.IRepository) *ReadAllUsersHandler {
	return &ReadAllUsersHandler{repo: repository}
}

// Handle read all users logic
func (h *ReadAllUsersHandler) Handle(ctx context.Context) ([]*user.User, error) {

	users, err := h.repo.ReadAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
