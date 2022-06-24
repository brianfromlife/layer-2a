package command

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// DeleteUserHandler class for all required interfaces to delete user
type DeleteUserHandler struct {
	repo user.IRepository
}

// NewDeleteUserHandler create new handler
func NewDeleteUserHandler(repository user.IRepository) *DeleteUserHandler {
	return &DeleteUserHandler{repo: repository}
}

// Handle delete user logic
func (h *DeleteUserHandler) Handle(ctx context.Context, ID int64) error {

	err := h.repo.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
