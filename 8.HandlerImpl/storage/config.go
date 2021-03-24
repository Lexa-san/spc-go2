package storage

type Config struct {
	// Строка подключения к БД
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
