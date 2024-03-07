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
	Redis          RedisConfig    `toml:"redis"`
}

type PostgresConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Database string `toml:"database"`
}

type RedisConfig struct {
	Host             string `toml:"host"`
	Port             int    `toml:"port"`
	Password         string `toml:"password"`
	DB               int    `toml:"db"`
	DisableIndentity bool   `toml:"disable_indentity"` // Disable set-info on connect
	DialTimeout      int    `toml:"dial_timeout"`
	MaxRetries       int    `toml:"max_retries"`
}
