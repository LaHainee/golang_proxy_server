package config

type AppConfig struct {
	LoggingLevel string       `toml:"logging_level"`
	ServerConfig ServerConfig `toml:"server"`
}

func NewConfig() *AppConfig {
	return &AppConfig{}
}

type ServerConfig struct {
	Port string `toml:"port"`
}
