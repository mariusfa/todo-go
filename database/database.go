package database

import (
	"fmt"
	"todo/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(dbConfig config.DbConfig, path string) error {
	connectionString := dbConfig.GetConnectionString()

	basePath := fmt.Sprintf("file://%s/base", path)
	m, err := migrate.New(basePath, connectionString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}