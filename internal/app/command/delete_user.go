package command

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// DeleteUserHandler
type DeleteUserHandler struct {
	repo user.IRepository
}

// NewCreateUserHandler
func NewDeleteUserHandler(repository user.IRepository) *DeleteUserHandler {
	return &DeleteUserHandler{repo: repository}
}

// Handle
func (h *DeleteUserHandler) Handle(ctx context.Context, ID int64) error {

	err := h.repo.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
