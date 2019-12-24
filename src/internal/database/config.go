package database

import "os"

type Config struct {
	DSN        string
	Database   string
}

func NewConfig() *Config {
	
	return &Config{
		DSN:      os.Getenv("DB_MONGO_DSN"),
		Database: os.Getenv("DB_MONGO_DATABASE"),
	}
}
