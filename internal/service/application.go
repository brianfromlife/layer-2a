package service

import (
	"context"

	"github.com/trevatk/layer-2a/internal/adapter"
	"github.com/trevatk/layer-2a/internal/app"
	"github.com/trevatk/layer-2a/internal/app/command"
	"github.com/trevatk/layer-2a/internal/app/query"
	"github.com/trevatk/layer-2a/internal/domain/user"
	"github.com/trevatk/layer-2a/internal/setup"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProvideApplication initialize all services to fulfill all interfaces
func ProvideApplication(lc fx.Lifecycle, log *zap.Logger, cfg *setup.Config) (*app.Application, error) {

	db, err := cfg.Database.ProvideDatabase(log)
	if err != nil {
		return nil, err
	}

	fc := &user.FactoryConfig{}
	factory := user.NewFactory(fc)

	repo := adapter.NewRepository(db, factory)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.Ping(ctx)
			},
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		},
	)

	return &app.Application{
		Commands: app.Commands{
			CreateUserHandler: command.NewCreateUserHandler(repo),
			UpdateUserHandler: command.NewUpdateUserHandler(repo),
			DeleteUserHandler: command.NewDeleteUserHandler(repo),
		},
		Queries: app.Queries{
			ReadUserHandler:     query.NewReadUserHandler(repo),
			ReadAllUsersHandler: query.NewReadAllUsersHandler(repo),
		},
	}, nil
}
