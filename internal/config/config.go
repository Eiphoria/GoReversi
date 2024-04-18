package config

import "os"

type Config struct {
	DBConf   DBConfig
	HashSalt string
	// Debug    bool
	// Port     int
	// Username string
	// Password string
	// Adress   string
	// DBName   string
}

type DBConfig struct {
	ConnectionURL string
}

func New() *Config {
	return &Config{
		DBConf: DBConfig{
			ConnectionURL: os.Getenv("DB_CONN"),
		},
		HashSalt: os.Getenv("HASH_SALT"),
	}
}
