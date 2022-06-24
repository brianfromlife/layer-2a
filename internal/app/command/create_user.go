package command

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/trevatk/layer-2a/internal/domain/user"
)

// CreateUserHandler class for all required interfaces to create user
type CreateUserHandler struct {
	repo user.IRepository
}

// NewCreateUserHandler create new instance
func NewCreateUserHandler(repository user.IRepository) *CreateUserHandler {
	return &CreateUserHandler{repo: repository}
}

// Handle create user logic
func (h *CreateUserHandler) Handle(ctx context.Context, u *user.User) (*user.User, error) {

	hashed, err := h.hashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashed

	user, err := h.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *CreateUserHandler) hashPassword(input string) (string, error) {

	s := sha256.New()
	_, err := s.Write([]byte(input))
	if err != nil {
		return "", fmt.Errorf("s.Write: %v", err)
	}

	return string(s.Sum(nil)), nil
}
