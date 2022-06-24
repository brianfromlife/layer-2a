package port

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/trevatk/layer-2a/internal/app"
	"github.com/trevatk/layer-2a/internal/transform"
	pb "github.com/trevatk/layer-2a/proto/auth_v1"
)

// AuthServer implementation of grpc auth interface
type AuthServer struct {
	pb.UnimplementedAuth_V1Server
	app    *app.Application
	logger *zap.SugaredLogger
}

// ProvideAuthServer configures and returns a new auth server
func ProvideAuthServer(log *zap.Logger, application *app.Application) *AuthServer {

	logger := log.Named("auth server").Sugar()

	return &AuthServer{app: application, logger: logger}
}

// CreateUser transform and create new user
func (as *AuthServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {

	as.logger.Debug("create user handler")

	user, err := as.app.Commands.CreateUserHandler.Handle(ctx, transform.NewUser(in))
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return transform.User(user), nil
}

// GetUser retrieve an individual user
func (as *AuthServer) GetUser(ctx context.Context, in *pb.UserLookup) (*pb.User, error) {

	as.logger.Debug("get user handler")

	user, err := as.app.Queries.ReadUserHandler.Handle(ctx, in.Id)
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return transform.User(user), nil
}

// GetAllUsers retrieve all users
func (as *AuthServer) GetAllUsers(ctx context.Context, empty *emptypb.Empty) (*pb.Users, error) {

	users, err := as.app.Queries.ReadAllUsersHandler.Handle(ctx)
	if err != nil {
		return nil, err
	}

	return transform.Users(users), nil
}

// UpdateUser transform and pass object to handler
func (as *AuthServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {

	as.logger.Debug("update user handler")

	user, err := as.app.Commands.UpdateUserHandler.Handle(ctx, transform.InUser(in))
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return transform.User(user), nil
}

// DeleteUser call to delete user handler
func (as *AuthServer) DeleteUser(ctx context.Context, in *pb.UserLookup) (*emptypb.Empty, error) {

	as.logger.Debug("delete user handler")

	err := as.app.Commands.DeleteUserHandler.Handle(ctx, in.Id)
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
