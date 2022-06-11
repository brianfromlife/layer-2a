package setup

import (
	"github.com/trevatk/common/database"
	"github.com/trevatk/common/server"
)

// Config
type Config struct {
	Database database.Config
	Server   server.Config
}

// ProvideConfig
func ProvideConfig() *Config {
	return &Config{}
}
