package config

import "os"

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetDBConfig() *DbConfig {
	return &DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}
