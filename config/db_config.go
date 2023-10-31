package config

import "fmt"

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (config *DbConfig) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DbName)
}
