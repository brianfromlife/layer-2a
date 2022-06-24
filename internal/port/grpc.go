package port

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/trevatk/layer-2a/internal/setup"
	pb "github.com/trevatk/layer-2a/proto/auth_v1"
)

// InvokeServer config and start grpc server
func InvokeServer(lc fx.Lifecycle, log *zap.Logger, cfg *setup.Config, as *AuthServer) error {

	s, err := cfg.Server.ProvideGrpcServer()
	if err != nil {
		return err
	}

	pb.RegisterAuth_V1Server(s.GRPCServer(), as)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go s.ServeGRPC()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				go s.GracefulShutdown(nil)
				return nil
			},
		},
	)

	return nil
}
