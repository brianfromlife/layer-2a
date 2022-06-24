package main

import (
	"context"
	"log"
	"time"

	"github.com/trevatk/common/logging"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/trevatk/layer-2a/internal/port"
	"github.com/trevatk/layer-2a/internal/service"
	"github.com/trevatk/layer-2a/internal/setup"
)

func main() {

	app := fx.New(
		fx.Provide(logging.ProvideLogger),
		fx.Provide(setup.ProvideConfig),
		fx.Invoke(setup.InvokeConfig),
		fx.Provide(service.ProvideApplication),
		fx.Provide(port.ProvideAuthServer),
		fx.Invoke(port.InvokeServer),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	start, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Start(start); err != nil {
		log.Fatal(err)
	}

	<-app.Done()

	stop, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Stop(stop); err != nil {
		log.Fatal(err)
	}
}
