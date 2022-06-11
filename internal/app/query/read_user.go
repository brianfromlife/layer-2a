package query

import (
	"context"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// ReadUserHandler
type ReadUserHandler struct {
	repo user.IRepository
}

// NewReadUserHandler
func NewReadUserHandler(repository user.IRepository) *ReadUserHandler {
	return &ReadUserHandler{repo: repository}
}

// Handle
func (h *ReadUserHandler) Handle(ctx context.Context, ID int64) (*user.User, error) {

	user, err := h.repo.ReadUser(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
