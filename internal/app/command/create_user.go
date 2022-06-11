package command

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// CreateUserHandler
type CreateUserHandler struct {
	repo user.IRepository
}

// NewCreateUserHandler
func NewCreateUserHandler(repository user.IRepository) *CreateUserHandler {
	return &CreateUserHandler{repo: repository}
}

// Handle
func (h *CreateUserHandler) Handle(ctx context.Context, u *user.User) (*user.User, error) {

	user, err := h.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
