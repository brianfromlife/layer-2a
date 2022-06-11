package app

import (
	"github.com/trevatk/layer-2a/internal/app/command"
	"github.com/trevatk/layer-2a/internal/app/query"
)

// Application
type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands
type Commands struct {
	CreateUserHandler *command.CreateUserHandler
	UpdateUserHandler *command.UpdateUserHandler
	DeleteUserHandler *command.DeleteUserHandler
}

// Queries
type Queries struct {
	ReadUserHandler     *query.ReadUserHandler
	ReadAllUsersHandler *query.ReadAllUsersHandler
}
