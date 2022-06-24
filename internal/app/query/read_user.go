package query

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// ReadUserHandler class for all required interfaces to read user
type ReadUserHandler struct {
	repo user.IRepository
}

// NewReadUserHandler create new instance
func NewReadUserHandler(repository user.IRepository) *ReadUserHandler {
	return &ReadUserHandler{repo: repository}
}

// Handle read user logic
func (h *ReadUserHandler) Handle(ctx context.Context, ID int64) (*user.User, error) {

	user, err := h.repo.ReadUser(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
