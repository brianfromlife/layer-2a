package command

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// UpdateUserHandler
type UpdateUserHandler struct {
	repo user.IRepository
}

// NewCreateUserHandler
func NewUpdateUserHandler(repository user.IRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repository}
}

// Handle
func (h *UpdateUserHandler) Handle(ctx context.Context, u *user.User) (*user.User, error) {

	user, err := h.repo.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
