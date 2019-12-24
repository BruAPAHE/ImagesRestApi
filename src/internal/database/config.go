package database

type Config struct {
	DSN        string
	Collection string
	Database   string
}

func NewConfig() *Config {
	return &Config{
		DSN:        "",
		Collection: "",
		Database:   "",
	}
}
