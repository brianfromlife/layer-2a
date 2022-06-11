package query

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// ReadAllUsersHandler
type ReadAllUsersHandler struct {
	repo user.IRepository
}

// NewReadAllUsersHandler
func NewReadAllUsersHandler(repository user.IRepository) *ReadAllUsersHandler {
	return &ReadAllUsersHandler{repo: repository}
}

// Handle
func (h *ReadAllUsersHandler) Handle(ctx context.Context) ([]*user.User, error) {

	users, err := h.repo.ReadAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
