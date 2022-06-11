package setup

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"
)

// InvokeConfig
func InvokeConfig(config *Config) error {
	return envconfig.ProcessWith(context.Background(), config, envconfig.OsLookuper())
}
