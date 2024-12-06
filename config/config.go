package config

import (
	"os"
)

type Configs interface{
	GetConfig() (*Config, error)
}

type Config struct{
	JWTSecret string

}
func GetConfig() (*Config, error) {
	return &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}