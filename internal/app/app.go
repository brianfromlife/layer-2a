package app

import (
	"github.com/trevatk/layer-2a/internal/app/command"
	"github.com/trevatk/layer-2a/internal/app/query"
)

// Application class for all commands and queries
type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands class to hold all commands
type Commands struct {
	CreateUserHandler *command.CreateUserHandler
	UpdateUserHandler *command.UpdateUserHandler
	DeleteUserHandler *command.DeleteUserHandler
}

// Queries class to hold all queries
type Queries struct {
	ReadUserHandler     *query.ReadUserHandler
	ReadAllUsersHandler *query.ReadAllUsersHandler
}
