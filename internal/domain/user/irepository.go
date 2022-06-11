package user

import "context"

// IRepository
type IRepository interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	ReadUser(ctx context.Context, ID int64) (*User, error)
	ReadAllUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, ID int64) error
}
