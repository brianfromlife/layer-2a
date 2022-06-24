package setup

import (
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
