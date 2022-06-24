package setup

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"
	"github.com/trevatk/common/database"
	"github.com/trevatk/common/server"
)

// Config environmental variables placeholder
type Config struct {
	Database database.Config
	Server   server.Config
}

// ProvideConfig create new config object
func ProvideConfig() *Config {
	return &Config{}
}

// InvokeConfig process all environmental variables
func InvokeConfig(config *Config) error {
	return envconfig.ProcessWith(context.Background(), config, envconfig.OsLookuper())
}
