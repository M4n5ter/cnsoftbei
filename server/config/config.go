package config

import (
	"os"

	"github.com/m4n5ter/cnsoftbei/common/log"
	"github.com/pelletier/go-toml/v2"
)

// Load reads the configuration from the given file path.
func Load(path string) (Config, error) {
	var c Config
	data, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}

	if err := toml.Unmarshal(data, &c); err != nil {
		return c, err
	}
	return c, nil
}

func MustLoad(path string) Config {
	c, err := Load(path)
	if err != nil {
		log.Panicf("failed to load configuration: %v", err)
	}
	return c
}

type Config struct {
	Host           string         `toml:"host"`
	Port           int            `toml:"port"`
	TrustedProxies []string       `toml:"trusted_proxies"`
	Postgres       PostgresConfig `toml:"postgres"`
}

type PostgresConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Database string `toml:"database"`
}
