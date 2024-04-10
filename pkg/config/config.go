package config

import (
	"os"

	"github.com/m4n5ter/cnsoftbei/pkg/yalog"
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

	return c, ParseTag("toml", c)
}

func MustLoad(path string) Config {
	c, err := Load(path)
	if err != nil {
		yalog.Fatalf("failed to load configuration: %v", err)
	}
	return c
}

type Config struct {
	Host           string         `toml:"host"`
	Port           int            `toml:"port,required"`
	TrustedProxies []string       `toml:"trusted_proxies"`
	Postgres       PostgresConfig `toml:"postgres,required"`
	Redis          RedisConfig    `toml:"redis,required"`
}

type PostgresConfig struct {
	Host     string `toml:"host,required"`
	Port     int    `toml:"port,required"`
	User     string `toml:"user,required"`
	Password string `toml:"password,required"`
	Schema   string `toml:"schema"`
	Database string `toml:"database,required"`
}

type RedisConfig struct {
	Host             string `toml:"host,required"`
	Port             int    `toml:"port,required"`
	Password         string `toml:"password"`
	DB               int    `toml:"db"`
	DisableIndentity bool   `toml:"disable_indentity"` // Disable set-info on connect
	DialTimeout      int    `toml:"dial_timeout"`
	MaxRetries       int    `toml:"max_retries"`
}
